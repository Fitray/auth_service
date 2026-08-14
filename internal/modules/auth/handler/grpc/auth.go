package auth_handler

import (
	"context"

	auth "github.com/Fitray/auth_service/api/gen/go"
	"github.com/Fitray/auth_service/internal/entities/dto"
	app_errors "github.com/Fitray/auth_service/internal/errors"
	"google.golang.org/grpc"
)

func Register(gRPCServer *grpc.Server, handler *AuthHandler) {
	auth.RegisterAuthServiceServer(gRPCServer, handler)
}

func (h *AuthHandler) Register(
	ctx context.Context,
	in *auth.RegisterRequest,
) (*auth.RegisterResponse, error) {
	input := dto.RegisterRequest{
		Login:       in.GetLogin(),
		Password:    in.GetPassword(),
		PhoneNumber: in.PhoneNumber,
	}

	registerResp, err := h.authService.RegisterUser(ctx, input)
	if err != nil {
		return nil, app_errors.ToGRPCError(err)
	}

	return &auth.RegisterResponse{
		UserId: registerResp.UserId,
	}, nil
}
