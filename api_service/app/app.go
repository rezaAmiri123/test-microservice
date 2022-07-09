package app

import (
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
)

type Application struct {
	Commands Commands
	//Queries Queries
}

type Commands struct {
	CreateUser    command.CreateUserHandler
	CreateArticle command.CreateArticleHandler
}

//type Queries struct {
//	GetProfile query.GetProfileHandler
//	GetUserToken query.GetUserTokenHandler
//}
