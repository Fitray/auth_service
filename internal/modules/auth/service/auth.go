package auth_service

import (
	"context"

	"github.com/Fitray/auth_service/internal/entities/dto"
)

func (s *AuthService) RegisterUser(
	ctx context.Context,
	registerReq dto.RegisterRequest,
) (dto.RegisterResponse, error) {
	userId, err := s.authPostgres.RegisterUser(ctx, registerReq)
	if err != nil {
		return dto.RegisterResponse{}, err
	}
	return dto.RegisterResponse{
		UserId: userId,
	}, nil
}
