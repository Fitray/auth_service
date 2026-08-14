package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/Fitray/auth_service/internal/config"
	"github.com/Fitray/auth_service/internal/logger"
	auth_module "github.com/Fitray/auth_service/internal/modules/auth"
	authorization_module "github.com/Fitray/auth_service/internal/modules/authorization"
	"github.com/Fitray/auth_service/internal/postgres"
	"github.com/Fitray/auth_service/internal/redis"
	grpc_server "github.com/Fitray/auth_service/internal/server/grpc"
	"google.golang.org/grpc"
)

func main() {
	shutdown_ctx, shutdown_cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer shutdown_cancel()

	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	appLogger, err := logger.NewLogger(
		config.LoggerConfig,
		config.ProjectConfig.ProjectRootPath,
	)
	if err != nil {
		panic(err)
	}
	defer appLogger.Close()

	postgresConn, err := postgres.NewConnectionPool(config.PostgresConfig)
	if err != nil {
		panic(fmt.Errorf("failed to connect to postgres: %w", err))
	}
	defer postgresConn.Close()

	redisClient, err := redis.NewClient(config.RedisConfig)
	if err != nil {
		panic(fmt.Errorf("failed to connect to redis: %w", err))
	}
	defer redisClient.Close()

	authHandler := auth_module.NewAuthModule(postgresConn, redisClient)
	authorizationHandler := authorization_module.NewAuthorizationModule(
		postgresConn,
		redisClient,
	)

	server := grpc_server.NewGRPCServer(
		authHandler,
		authorizationHandler,
		appLogger.Logger(),
		config.GRPCConfig,
	)

	go func() {
		<-shutdown_ctx.Done()

		ctx, cancel := context.WithTimeout(
			context.Background(),
			config.GRPCConfig.ShutdownTimeout,
		)
		defer cancel()

		server.Shutdown(ctx)
		postgresConn.Close()
		redisClient.Close()
	}()

	if err := server.Run(config.GRPCConfig, shutdown_ctx); err != nil {
		if !errors.Is(err, grpc.ErrServerStopped) {
			appLogger.Logger().Error(
				err.Error(),
			)
			return
		}
	}
}
