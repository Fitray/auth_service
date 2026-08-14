package auth_module

import (
	auth_handler "github.com/Fitray/auth_service/internal/modules/auth/handler/grpc"
	auth_postgres "github.com/Fitray/auth_service/internal/modules/auth/repository/postgres"
	auth_redis "github.com/Fitray/auth_service/internal/modules/auth/repository/redis"
	auth_service "github.com/Fitray/auth_service/internal/modules/auth/service"
	"github.com/Fitray/auth_service/internal/postgres"
	"github.com/Fitray/auth_service/internal/redis"
)

func NewAuthModule(
	pool postgres.Pool,
	redisClient redis.Redis,
) *auth_handler.AuthHandler {
	authPostgres := auth_postgres.NewAuthPostgres(pool)
	authRedis := auth_redis.NewAuthRedis(redisClient)
	authService := auth_service.NewAuthService(authPostgres, authRedis)

	return auth_handler.NewAuthHandler(authService)
}
