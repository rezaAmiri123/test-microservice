package agent

import (
	"github.com/rezaAmiri123/test-microservice/message_service/adapters/pg"
	"github.com/rezaAmiri123/test-microservice/message_service/app"
	"github.com/rezaAmiri123/test-microservice/message_service/app/commands"
	"github.com/rezaAmiri123/test-microservice/message_service/app/queries"
	"github.com/rezaAmiri123/test-microservice/pkg/db/postgres"
)

func (a *Agent) setupApplication() error {
	dbConn, err := postgres.NewPsqlDB(a.DBConfig)
	if err != nil {
		return err
	}

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	repo := pg.NewPGEmailRepository(dbConn)

	application := &app.Application{
		Commands: app.Commands{
			CreateEmail: commands.NewCreateUserHandler(repo),
		},
		Queries: app.Queries{
			GetEmailByUUID: queries.NewGetEmailHandler(repo),
		},
	}
	a.Application = application
	return nil
}
