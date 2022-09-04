//go:generate mockgen -source repository.go -destination mock/repository.go -package wallet_mock
package wallet

import (
	"context"
)

type Repository interface {
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	AddWalletBalance(ctx context.Context, arg AddWalletBalanceParams) (Wallet, error)
	CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	GetWallet(ctx context.Context, id int64) (Wallet, error)
	GetWalletForUpdate(ctx context.Context, id int64) (Wallet, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	ListWallets(ctx context.Context, arg ListWalletParams) ([]Wallet, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error)
}
