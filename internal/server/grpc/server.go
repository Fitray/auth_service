package grpc_server

import (
	"net"

	"github.com/Fitray/auth_service/internal/config"
	handler "github.com/Fitray/auth_service/internal/handler/grpc"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	Server *grpc.Server
}

func NewGRPCServer(serverApi *handler.ServerAPI) *GRPCServer {
	server := grpc.NewServer()
	handler.Register(server, serverApi)
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
