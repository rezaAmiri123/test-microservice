package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/domain/wallet"
)

const createWallet = `INSERT INTO wallets (owner, balance, currency) VALUES ($1, $2, $3) returns 
						(wallet_id, owner, balance, currency, created_at, updated_at)`

func (r *PGWalletRepository) CreateWallet(ctx context.Context, arg wallet.CreateWalletParams) (res wallet.Wallet, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGWalletRepository.Create")
	defer span.Finish()

	if err = r.DB.QueryRowxContext(
		ctx,
		createWallet,
		arg.Owner,
		arg.Balance,
		arg.Currency,

		// ).Scan(&res); err != nil {
	).StructScan(&res); err != nil {
		return res, fmt.Errorf("connot create wallet: %w", err)
	}
	return
}
