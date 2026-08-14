package grpc_server

import (
	"context"
	"fmt"
	"net"

	"github.com/Fitray/auth_service/internal/config"
	auth_handler "github.com/Fitray/auth_service/internal/modules/auth/handler/grpc"
	authorization_handler "github.com/Fitray/auth_service/internal/modules/authorization/handler/grpc"
	"github.com/Fitray/auth_service/internal/server/grpc/interceptor"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Server *grpc.Server
}

func NewGRPCServer(
	authHandler *auth_handler.AuthHandler,
	authorizationHandler *authorization_handler.AuthorizationHandler,
	log *zap.Logger,
	config config.GRPCConfig,
) *GRPCServer {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.RequestContextInterceptor(),
			interceptor.TimeoutInterceptor(config.Timeout),
			interceptor.LoggingUnaryInterceptor(log),
			interceptor.RecoveryUnaryInterceptor(log),
		),
	)
	auth_handler.Register(server, authHandler)
	authorization_handler.Register(server, authorizationHandler)
	reflection.Register(server)
	return &GRPCServer{
		Server: server,
	}
}

func (s *GRPCServer) Run(
	cfg config.GRPCConfig,
	shutdown_ctx context.Context,
) error {
	l, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return err
	}

	fmt.Println("Starting server", cfg.Port)
	if err := s.Server.Serve(l); err != nil {
		return err
	}
	return nil
}

func (s *GRPCServer) Shutdown(ctx context.Context) {
	gracefulDone := make(chan struct{})

	go func() {
		s.Server.GracefulStop()
		close(gracefulDone)
	}()

	select {
	case <-gracefulDone:
		fmt.Println("Server stopped")
	case <-ctx.Done():
		s.Server.Stop()
	}
}
