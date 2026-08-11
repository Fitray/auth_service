package main

import (
	"fmt"

	"github.com/Fitray/auth_service/internal/config"
	handler "github.com/Fitray/auth_service/internal/handler/grpc"
	"github.com/Fitray/auth_service/internal/logger"
	"github.com/Fitray/auth_service/internal/postgres"
	"github.com/Fitray/auth_service/internal/redis"
	"github.com/Fitray/auth_service/internal/repository"
	grpc_server "github.com/Fitray/auth_service/internal/server/grpc"
	"github.com/Fitray/auth_service/internal/service"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger(
		config.LoggerConfig,
		config.ProjectConfig.ProjectRootPath,
	)
	if err != nil {
		panic(err)
	}

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

	grpcRepo := repository.NewRepository()
	grpcService := service.NewService(grpcRepo)
	serverApi := handler.NewServerAPI(grpcService)

	server := grpc_server.NewGRPCServer(serverApi)
	if err := server.Run(config.GRPCConfig); err != nil {
		fmt.Println(err)
	}
}
