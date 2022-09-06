package agent

import (
	"fmt"

	"github.com/rezaAmiri123/test-microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/adapters/pg"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/app"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/app/commands"
)

func (a *Agent) setupApplication() error {
	dbConn, err := postgres.NewPsqlDB(a.DBConfig)
	if err != nil {
		return fmt.Errorf("cannot load db: %w", err)
	}

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	repo := pg.NewPGWalletRepository(dbConn)

	application := &app.Application{
		Commands: app.Commands{
			CreateWallet: commands.NewCreateWalletHandler(repo),
		},
		Queries: app.Queries{},
	}
	a.Application = application
	return nil
}
