package authorization_module

import (
	authorization_handler "github.com/Fitray/auth_service/internal/modules/authorization/handler/grpc"
	authorization_postgres "github.com/Fitray/auth_service/internal/modules/authorization/repository/postgres"
	authorization_redis "github.com/Fitray/auth_service/internal/modules/authorization/repository/redis"
	authorization_service "github.com/Fitray/auth_service/internal/modules/authorization/service"
	"github.com/Fitray/auth_service/internal/postgres"
	"github.com/Fitray/auth_service/internal/redis"
)

func NewAuthorizationModule(
	pool postgres.Pool,
	redisClient redis.Redis,
) *authorization_handler.AuthorizationHandler {
	authorizationPostgres := authorization_postgres.NewAuthorizationPostgres(pool)
	authorizationRedis := authorization_redis.NewAuthorizationRedis(redisClient)
	authorizationService := authorization_service.NewAuthorizationService(
		authorizationPostgres,
		authorizationRedis,
	)

	return authorization_handler.NewAuthorizationHandler(authorizationService)
}
