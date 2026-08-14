package auth_redis

import "github.com/Fitray/auth_service/internal/redis"

type AuthRedis struct {
	client redis.Redis
}

func NewAuthRedis(client redis.Redis) *AuthRedis {
	return &AuthRedis{
		client: client,
	}
}
