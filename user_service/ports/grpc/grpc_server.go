package grpc

import (
	"github.com/rezaAmiri123/test-microservice/user_service/app"
	userservice "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
	"google.golang.org/grpc"
)

type Config struct {
	App *app.Application
}

var _ userservice.UsersServiceServer = (*grpcServer)(nil)

type grpcServer struct {
	cfg *Config
	userservice.UnimplementedUsersServiceServer
}

func newGRPCServer(config *Config) (*grpcServer, error) {
	srv := &grpcServer{
		cfg: config,
	}
	return srv, nil
}

func NewGRPCServer(config *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
	gsrv := grpc.NewServer(opts...)
	srv, err := newGRPCServer(config)
	if err != nil {
		return nil, err
	}
	userservice.RegisterUsersServiceServer(gsrv, srv)
	return gsrv, nil
}
