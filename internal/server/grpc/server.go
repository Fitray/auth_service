package grpc_server

import (
	"net"

	"github.com/Fitray/auth_service/internal/config"
	handler "github.com/Fitray/auth_service/internal/handler/grpc"
	"github.com/Fitray/auth_service/internal/server/grpc/interceptor"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	Server *grpc.Server
}

func NewGRPCServer(
	serverApi *handler.ServerAPI,
	log *zap.Logger,
	config config.GRPCConfig,
) *GRPCServer {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.RequestContextInterceptor(),
			interceptor.TimeoutInterceptor(config.Timeout),
			interceptor.LoggingUnaryInterceptor(log),
		),
	)
	handler.Register(server, serverApi)
	reflection.Register(server)
	return &GRPCServer{
		Server: server,
	}
}

func (s *GRPCServer) Run(
	cfg config.GRPCConfig,
) error {
	l, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return err
	}

	if err := s.Server.Serve(l); err != nil {
		return err
	}
	return nil
}

func (s *GRPCServer) Stop() {
	s.Server.GracefulStop()
}
