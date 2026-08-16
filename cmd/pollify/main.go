package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	core_config "pollify/internal/core/config"
	core_logger "pollify/internal/core/logger"
	core_pgx_pool "pollify/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "pollify/internal/core/transport/http/middleware"
	core_http_server "pollify/internal/core/transport/http/server"
	users_postgres_repository "pollify/internal/features/users/repository/postgres"
	users_service "pollify/internal/features/users/service"
	users_transport_http "pollify/internal/features/users/transport/http"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	config := core_config.NewConfigMust()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Printf("Failed to initialize logger: %s\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initialize postgres connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, config)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository, config.JWTToken)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouter := core_http_server.NewAPIVersionRouter(&core_http_server.APIVersion1)

	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
