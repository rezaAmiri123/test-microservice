package wallet

import "github.com/google/uuid"

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromWalletID int64 `json:"from_wallet_id"`
	ToWalletID   int64 `json:"to_wallet_id"`
	Amount       int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer   Transfer `json:"transfer"`
	FromWallet Wallet   `json:"from_wallet"`
	ToWallet   Wallet   `json:"to_wallet"`
	FromEntry  Entry    `json:"from_entry"`
	ToEntry    Entry    `json:"to_entry"`
}

type CreateWalletParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type AddWalletBalanceParams struct {
	ID     uuid.UUID `json:"id"`
	Amount int64     `json:"amount"`
}

type ListWalletParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type UpdateWalletParams struct {
	ID      uuid.UUID `json:"id"`
	Balance int64     `json:"balance"`
}

type CreateEntryParams struct {
	WalletID uuid.UUID `json:"wallet_id"`
	Amount   int64     `json:"amount"`
}

type ListEntriesParams struct {
	WalletID uuid.UUID `json:"wallet_id"`
	Limit    int32     `json:"limit"`
	Offset   int32     `json:"offset"`
}

type CreateTransferParams struct {
	FromWalletID uuid.UUID `json:"from_wallet_id"`
	ToWalletID   uuid.UUID `json:"to_wallet_id"`
	Amount       int64     `json:"amount"`
}
type ListTransfersParams struct {
	FromWalletID int64 `json:"from_wallet_id"`
	ToWalletID   int64 `json:"to_wallet_id"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}
