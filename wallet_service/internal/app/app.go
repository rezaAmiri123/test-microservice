package app

import "github.com/rezaAmiri123/test-microservice/wallet_service/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateWallet *commands.CreateWalletHandler
}
