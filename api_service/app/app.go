package app

import (
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
)

type Application struct {
	Commands Commands
	//Queries Queries
}

type Commands struct {
	CreateUser command.CreateUserHandler
}

//type Queries struct {
//	GetProfile query.GetProfileHandler
//	GetUserToken query.GetUserTokenHandler
//}
