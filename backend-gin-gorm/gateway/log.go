package gateway

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	slogotel "github.com/remychantenay/slog-otel"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

// OTLPLogConfig holds OTLP log exporter settings.
type OTLPLogConfig struct {
	Endpoint string `yaml:"endpoint" validate:"required"`
	Insecure bool   `yaml:"insecure"`
}

// UptraceLogConfig holds Uptrace log exporter settings.
type UptraceLogConfig struct {
	Endpoint string `yaml:"endpoint" validate:"required"`
	DSN      string `yaml:"dsn" validate:"required"`
}

// LogConfig holds logging configuration including level, platform, and exporter settings.
type LogConfig struct {
	Level    string            `yaml:"level"`
	Platform string            `yaml:"platform"`
	Levels   map[string]string `yaml:"levels"`
	Exporter string            `yaml:"exporter" validate:"oneof=none otlphttp uptracehttp"`
	OTLP     *OTLPLogConfig    `yaml:"otlp"`
	Uptrace  *UptraceLogConfig `yaml:"uptrace"`
}

// InitLog sets up the global slog logger. If the exporter is "none", it uses a local JSON handler;
// otherwise it delegates to InitLogProvider for OTLP/Uptrace export.
func InitLog(ctx context.Context, logConfig *LogConfig, appName string) (func(), error) {
	if logConfig.Exporter == "none" {
		defaultLogLevel := stringToLogLevel(logConfig.Level)
		jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ //nolint:exhaustruct
			Level: defaultLogLevel,
		})
		handler := slogotel.OtelHandler{ //nolint:exhaustruct
			Next: jsonHandler,
		}
		slog.SetDefault(slog.New(handler))
		return func() {
			// No-op
		}, nil
	}

	return InitLogProvider(ctx, logConfig, appName)
}

const logShutdownTimeout = 5 * time.Second

// InitLogExporterFunc is a function type that creates a log exporter from config.
type InitLogExporterFunc func(ctx context.Context, logConfig *LogConfig) (sdklog.Exporter, error)

func initLogExporter(ctx context.Context, logConfig *LogConfig) (sdklog.Exporter, error) {
	initLogExporters := map[string]InitLogExporterFunc{
		"otlphttp":    initLogExporterOTLPHTTP,
		"uptracehttp": initLogExporterUptraceHTTP,
	}

	initLogExporter, ok := initLogExporters[logConfig.Exporter]
	if !ok {
		return nil, fmt.Errorf("invalid log exporter: %s", logConfig.Exporter)
	}

	return initLogExporter(ctx, logConfig)
}

// InitLogProvider creates an OpenTelemetry log provider with batch processing
// and returns a shutdown function.
func InitLogProvider(ctx context.Context, logConfig *LogConfig, appName string) (func(), error) {
	exp, err := initLogExporter(ctx, logConfig)
	if err != nil {
		return nil, fmt.Errorf("initLogExporter: %w", err)
	}

	bp := sdklog.NewBatchProcessor(exp,
		sdklog.WithMaxQueueSize(10_000),
		sdklog.WithExportMaxBatchSize(10_000),
		sdklog.WithExportInterval(10*time.Second),
		sdklog.WithExportTimeout(10*time.Second),
	)

	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(bp),
		sdklog.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appName),
		)),
	)

	global.SetLoggerProvider(lp)

	defaultLogLevel := stringToLogLevel(logConfig.Level)
	otelHandler := otelslog.NewHandler(appName, otelslog.WithLoggerProvider(lp))
	filteredHandler := &levelFilterHandler{
		handler:  otelHandler,
		minLevel: defaultLogLevel,
	}

	slog.SetDefault(slog.New(filteredHandler))

	return func() {
		shutdownBaseCtx := context.Background()
		if ctx != nil {
			shutdownBaseCtx = context.WithoutCancel(ctx)
		}
		shutdownCtx, cancel := context.WithTimeout(shutdownBaseCtx, logShutdownTimeout)
		defer cancel()

		if err := lp.Shutdown(shutdownCtx); err != nil {
			slog.Error("failed to shutdown log provider", slog.Any("error", err))
		}
		if err := bp.Shutdown(shutdownCtx); err != nil {
			slog.Error("failed to shutdown log batch processor", slog.Any("error", err))
		}
	}, nil
}

type levelFilterHandler struct {
	handler  slog.Handler
	minLevel slog.Level
}

func (h *levelFilterHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.minLevel
}

func (h *levelFilterHandler) Handle(ctx context.Context, record slog.Record) error {
	if !h.Enabled(ctx, record.Level) {
		return nil
	}
	return h.handler.Handle(ctx, record) //nolint:wrapcheck
}

func (h *levelFilterHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &levelFilterHandler{
		handler:  h.handler.WithAttrs(attrs),
		minLevel: h.minLevel,
	}
}

func (h *levelFilterHandler) WithGroup(name string) slog.Handler {
	return &levelFilterHandler{
		handler:  h.handler.WithGroup(name),
		minLevel: h.minLevel,
	}
}

func stringToLogLevel(value string) slog.Level {
	switch strings.ToLower(value) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		slog.Info("Unsupported log level: " + value)

		return slog.LevelWarn
	}
}
