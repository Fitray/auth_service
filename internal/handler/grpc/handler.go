package handler

import (
	auth "github.com/Fitray/auth_service/api/gen/go"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	auth.UnimplementedAuthServiceServer
	service Service
}

type Service interface {
}

func NewServerAPI(
	service Service,
) *ServerAPI {
	return &ServerAPI{
		service: service,
	}
}

func Register(
	gRPC *grpc.Server,
	serverApi *ServerAPI,
) {
	auth.RegisterAuthServiceServer(
		gRPC,
		serverApi,
	)
}
