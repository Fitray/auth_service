package auth_service

import (
	"context"

	"github.com/Fitray/auth_service/internal/entities/dto"
)

type AuthService struct {
	authPostgres AuthPostgres
	authRedis    AuthRedis
}

type AuthPostgres interface {
	RegisterUser(
		ctx context.Context,
		registerReq dto.RegisterRequest,
	) (int64, error)
}
type AuthRedis interface{}

func NewAuthService(
	authPostgres AuthPostgres,
	authRedis AuthRedis,
) *AuthService {
	return &AuthService{
		authPostgres: authPostgres,
		authRedis:    authRedis,
	}
}
