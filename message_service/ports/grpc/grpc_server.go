package grpc

import (
	"github.com/rezaAmiri123/test-microservice/message_service/app"
	messageservice "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"google.golang.org/grpc"
)

type Config struct {
	//Metric *metrics.LibraryServiceMetric
	App *app.Application
}

var _ messageservice.MessageServiceServer = (*MessageGRPCServer)(nil)

type MessageGRPCServer struct {
	cfg *Config
	messageservice.UnimplementedMessageServiceServer
}

func NewMessageGRPCServer(config *Config) (*MessageGRPCServer, error) {
	srv := &MessageGRPCServer{
		cfg: config,
	}
	return srv, nil
}

func NewGRPCServer(config *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
	gsrv := grpc.NewServer(opts...)
	srv, err := NewMessageGRPCServer(config)
	if err != nil {
		return nil, err
	}
	messageservice.RegisterMessageServiceServer(gsrv, srv)
	return gsrv, nil
}
