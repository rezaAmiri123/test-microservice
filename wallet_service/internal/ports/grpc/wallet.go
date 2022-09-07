package grpc

import (
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/domain/wallet"
	walletapi "github.com/rezaAmiri123/test-microservice/wallet_service/proto/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func WalletResponseToGrpc(w *wallet.Wallet) *walletapi.Wallet {
	res := &walletapi.Wallet{}
	res.Balance = w.Balance
	res.Currency = w.Currency
	res.Owner = w.Owner
	res.WalletId = w.WalletID[:]
	res.CreatedAt = timestamppb.New(w.CreatedAt)
	res.UpdatedAt = timestamppb.New(w.UpdatedAt)
	return res
}
