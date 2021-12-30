package agent

import (
	"github.com/rezaAmiri123/test-microservice/user_service/adapters"
	"github.com/rezaAmiri123/test-microservice/user_service/app"
	"github.com/rezaAmiri123/test-microservice/user_service/app/command"
	"github.com/rezaAmiri123/test-microservice/user_service/app/query"
)

func (a *Agent) setupApplication() error {
	repo, err := adapters.NewGORMUserRepository(a.DBConfig)
	if err != nil {
		return err
	}
	application := &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(repo),
		},
		Queries: app.Queries{
			GetProfile: query.NewGetProfileHandler(repo),
			GetUserToken: query.NewGetUserTokenHandler(repo),
		},
	}
	a.Application = application
	return nil
}
