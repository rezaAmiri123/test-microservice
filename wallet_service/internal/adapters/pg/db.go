package pg

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/domain/wallet"
)

func NewPGWalletRepository(db *sqlx.DB) *PGWalletRepository {
	return &PGWalletRepository{DB: db}
}

// News Repository
type PGWalletRepository struct {
	DB *sqlx.DB
}

func (r *PGWalletRepository) TransferTx(ctx context.Context, arg wallet.TransferTxParams) (wallet.TransferTxResult, error) {
	return wallet.TransferTxResult{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) AddWalletBalance(ctx context.Context, arg wallet.AddWalletBalanceParams) (wallet.Wallet, error) {
	return wallet.Wallet{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) CreateEntry(ctx context.Context, arg wallet.CreateEntryParams) (wallet.Entry, error) {
	return wallet.Entry{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) CreateTransfer(ctx context.Context, arg wallet.CreateTransferParams) (wallet.Transfer, error) {
	return wallet.Transfer{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) GetWallet(ctx context.Context, id int64) (wallet.Wallet, error) {
	return wallet.Wallet{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) GetWalletForUpdate(ctx context.Context, id int64) (wallet.Wallet, error) {
	return wallet.Wallet{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) GetEntry(ctx context.Context, id int64) (wallet.Entry, error) {
	return wallet.Entry{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) GetTransfer(ctx context.Context, id int64) (wallet.Transfer, error) {
	return wallet.Transfer{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) ListWallets(ctx context.Context, arg wallet.ListWalletParams) ([]wallet.Wallet, error) {
	return []wallet.Wallet{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) ListEntries(ctx context.Context, arg wallet.ListEntriesParams) ([]wallet.Entry, error) {
	return []wallet.Entry{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) ListTransfers(ctx context.Context, arg wallet.ListTransfersParams) ([]wallet.Transfer, error) {
	return []wallet.Transfer{}, fmt.Errorf("not emplimented")
}
func (r *PGWalletRepository) UpdateWallet(ctx context.Context, arg wallet.UpdateWalletParams) (wallet.Wallet, error) {
	return wallet.Wallet{}, fmt.Errorf("not emplimented")
}
