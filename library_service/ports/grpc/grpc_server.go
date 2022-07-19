package grpc

import (
	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/metrics"
	libraryservice "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"google.golang.org/grpc"
)

type Config struct {
	Metric *metrics.LibraryServiceMetric
	App    *app.Application
}

var _ libraryservice.ArticleServiceServer = (*ArticleGRPCServer)(nil)

type ArticleGRPCServer struct {
	cfg *Config
	libraryservice.UnimplementedArticleServiceServer
}

func NewArticleGRPCServer(config *Config) (*ArticleGRPCServer, error) {
	srv := &ArticleGRPCServer{
		cfg: config,
	}
	return srv, nil
}

func NewGRPCServer(config *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
	gsrv := grpc.NewServer(opts...)
	srv, err := NewArticleGRPCServer(config)
	if err != nil {
		return nil, err
	}
	libraryservice.RegisterArticleServiceServer(gsrv, srv)
	return gsrv, nil
}
