package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/config"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/handler"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/middleware"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/gateway"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/process"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/usecase"
)

func main() {
	exitCode, err := run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(exitCode)
}

func run() (int, error) {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		return 1, fmt.Errorf("load config: %w", err)
	}
	logger := slog.Default().With(slog.String(domain.LoggerNameKey, domain.AppName+"-main"))

	dbc, shutdownDB, err := gateway.InitDB(ctx, cfg.DB, cfg.Log, domain.AppName)
	if err != nil {
		return 1, fmt.Errorf("init db: %w", err)
	}
	defer shutdownDB()
	router, err := handler.InitRootRouterGroup(ctx, cfg.Server.Gin, domain.AppName)
	if err != nil {
		return 1, fmt.Errorf("init router: %w", err)
	}

	authTokenManager := gateway.NewAuthTokenManager(
		[]byte(cfg.Auth.SigningKey),
		jwt.SigningMethodHS256,
		time.Duration(cfg.Auth.AccessTokenTTLMin)*time.Minute,
		time.Duration(cfg.Auth.Cookie.RefreshThresholdMin)*time.Minute,
	)
	authUsecase := usecase.NewAuthUsecase(authTokenManager)

	// api
	api := router.Group("api")

	// v1
	v1 := api.Group("v1")

	authMiddleware := middleware.NewAuthMiddleware(authUsecase, cfg.Auth.Cookie, cfg.Auth.AccessTokenTTLMin)
	{
		todoRepo := gateway.NewTodoRepository(dbc.DB)
		todoCreateBulkCommandTxManager := gateway.NewTodoCreateBulkCommandTxManager(dbc)
		todoUsecase := usecase.NewTodoUsecase(todoRepo, todoCreateBulkCommandTxManager)
		funcs := handler.NewInitTodoRouterFunc(todoUsecase)
		funcs(v1, authMiddleware)
	}
	{
		funcs := handler.NewInitAuthRouterFunc(authUsecase, cfg.Auth.Cookie, cfg.Auth.AccessTokenTTLMin)
		funcs(v1)
	}

	// run
	readHeaderTimeout := time.Duration(cfg.Server.ReadHeaderTimeoutSec) * time.Second
	shutdownTime := time.Duration(cfg.Server.Shutdown.TimeSec1) * time.Second
	result := process.Run(ctx,
		controller.WithWebServerProcess(router, cfg.Server.HTTPPort, readHeaderTimeout, shutdownTime),
		controller.WithMetricsServerProcess(cfg.Server.MetricsPort, readHeaderTimeout, shutdownTime),
		gateway.WithSignalWatchProcess(),
	)

	gracefulShutdownTime2 := time.Duration(cfg.Server.Shutdown.TimeSec2) * time.Second
	time.Sleep(gracefulShutdownTime2)
	logger.InfoContext(ctx, "exited")
	return result, nil
}
