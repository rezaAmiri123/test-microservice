package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	walletservice "github.com/rezaAmiri123/test-microservice/wallet_service/proto/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *WalletGRPCServer) CreateUser(ctx context.Context, req *walletservice.CreateWalletRequest) (*walletservice.CreateWalletResponse, error) {
	//s.cfg.Metric.GetArticleBySlugGrpcRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(ctx, "MessageGRPCServer.CreateEmail")
	defer span.Finish()

	wallet, err := s.cfg.App.Commands.CreateWallet.Handle(ctx, req.GetCurrency(), req.GetOwner())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &walletservice.CreateWalletResponse{Wallet: wallet}, nil
}
