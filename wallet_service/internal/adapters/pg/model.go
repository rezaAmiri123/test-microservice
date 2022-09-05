package pg

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	WalletID  uuid.UUID `json:"wallet_id" db:"wallet_id"`
	Owner     string    `json:"owner" db:"owner"`
	Balance   int64     `json:"balance" db:"balance"`
	Currency  string    `json:"currency" db:"currency"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
