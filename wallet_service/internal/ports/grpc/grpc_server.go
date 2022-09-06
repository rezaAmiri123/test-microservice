package grpc

import (
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/app"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/metrics"
	walletservice "github.com/rezaAmiri123/test-microservice/wallet_service/proto/grpc"
	"google.golang.org/grpc"
)

type Config struct {
	Metric *metrics.WalletServiceMetric
	App    *app.Application
}

var _ walletservice.WalletServiceServer = (*WalletGRPCServer)(nil)

type WalletGRPCServer struct {
	cfg *Config
	walletservice.UnimplementedWalletServiceServer
}

func NewWalletGRPCServer(config *Config) (*WalletGRPCServer, error) {
	srv := &WalletGRPCServer{
		cfg: config,
	}
	return srv, nil
}

func NewGRPCServer(config *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
	gsrv := grpc.NewServer(opts...)
	srv, err := NewWalletGRPCServer(config)
	if err != nil {
		return nil, err
	}
	walletservice.RegisterWalletServiceServer(gsrv, srv)
	return gsrv, nil
}
