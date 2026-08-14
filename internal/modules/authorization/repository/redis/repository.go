package authorization_redis

import "github.com/Fitray/auth_service/internal/redis"

type AuthorizationRedis struct {
	client redis.Redis
}

func NewAuthorizationRedis(client redis.Redis) *AuthorizationRedis {
	return &AuthorizationRedis{
		client: client,
	}
}
