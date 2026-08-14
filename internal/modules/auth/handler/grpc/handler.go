package auth_handler

import (
	"context"

	auth "github.com/Fitray/auth_service/api/gen/go"
	"github.com/Fitray/auth_service/internal/entities/dto"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	authService AuthService
}

type AuthService interface {
	RegisterUser(
		ctx context.Context,
		registerReq dto.RegisterRequest,
	) (dto.RegisterResponse, error)
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}
