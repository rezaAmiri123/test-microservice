package app

import (
	"github.com/rezaAmiri123/test-microservice/user_service/app/command"
	"github.com/rezaAmiri123/test-microservice/user_service/app/query"
)

type Application struct {
	Commands Commands
	Queries Queries
}

type Commands struct {
	CreateUser command.CreateUserHandler
}

type Queries struct {
	GetProfile query.GetProfileHandler
	GetUserToken query.GetUserTokenHandler
}
