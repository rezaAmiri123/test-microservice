package pg

import "github.com/jmoiron/sqlx"

func NewPGWalletRepository(db *sqlx.DB) *PGWalletRepository {
	return &PGWalletRepository{DB: db}
}

// News Repository
type PGWalletRepository struct {
	DB *sqlx.DB
}
