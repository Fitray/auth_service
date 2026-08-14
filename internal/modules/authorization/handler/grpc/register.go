package authorization_handler

import (
	auth "github.com/Fitray/auth_service/api/gen/go"
	"google.golang.org/grpc"
)

func Register(gRPCServer *grpc.Server, handler *AuthorizationHandler) {
	auth.RegisterAuthorizationServiceServer(gRPCServer, handler)
}
