package commands

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/domain/wallet"
)

type CreateWalletHandler struct {
	repo wallet.Repository
}

func NewCreateArticleHandler(repo wallet.Repository) *CreateWalletHandler {
	if repo == nil {
		panic("repo is nil")
	}
	return &CreateWalletHandler{repo: repo}
}

func (h CreateWalletHandler) Handle(ctx context.Context, currency, owner string) (*wallet.Wallet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateWalletHandler.Handle")
	defer span.Finish()

	arg := wallet.CreateWalletParams{
		Owner:    owner,
		Currency: currency,
		Balance:  0,
	}

	res, err := h.repo.CreateWallet(ctx, arg)
	return &res, err
}
