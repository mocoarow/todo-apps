// Package handler provides Gin HTTP handlers for the REST API,
// including request parsing, response mapping, routing, and CORS configuration.
package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/middleware"
)

// LogConfig controls access log output settings.
type LogConfig struct {
	AccessLog             bool `yaml:"accessLog"`
	AccessLogRequestBody  bool `yaml:"accessLogRequestBody"`
	AccessLogResponseBody bool `yaml:"accessLogResponseBody"`
}

// DebugConfig controls debug-mode features (Gin debug mode, artificial wait).
type DebugConfig struct {
	Gin  bool `yaml:"gin"`
	Wait bool `yaml:"wait"`
}

// Config holds handler-level configuration for CORS, logging, and debug settings.
type Config struct {
	CORS  *CORSConfig  `yaml:"cors" validate:"required"`
	Log   *LogConfig   `yaml:"log" validate:"required"`
	Debug *DebugConfig `yaml:"debug" validate:"required"`
}

// InitRootRouterGroup creates a Gin engine with recovery, CORS, metrics, tracing, and optional access logging.
func InitRootRouterGroup(_ context.Context, config *Config, appName string) *gin.Engine {
	if !config.Debug.Gin {
		gin.SetMode(gin.ReleaseMode)
	}

	corsConfig := InitCORS(config.CORS)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.New(corsConfig))
	router.Use(middleware.PrometheusMiddleware())
	router.Use(otelgin.Middleware(appName, otelgin.WithFilter(func(req *http.Request) bool {
		return req.URL.Path != "/"
	})))

	if config.Log.AccessLog {
		withRequestBody := false
		if config.Log.AccessLogRequestBody {
			withRequestBody = true
		}
		withResponseBody := false
		if config.Log.AccessLogResponseBody {
			withResponseBody = true
		}
		router.Use(sloggin.NewWithConfig(slog.Default(), sloggin.Config{ //nolint:exhaustruct
			DefaultLevel:     slog.LevelInfo,
			ClientErrorLevel: slog.LevelWarn,
			ServerErrorLevel: slog.LevelError,
			WithRequestBody:  withRequestBody,
			WithResponseBody: withResponseBody,
			Filters: []sloggin.Filter{
				func(c *gin.Context) bool {
					path := c.Request.URL.Path
					return path != "/"
				},
			},
		}))
	}

	if config.Debug.Wait {
		router.Use(middleware.NewWaitMiddleware(time.Second))
	}

	return router
}
