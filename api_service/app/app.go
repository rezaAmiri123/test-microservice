package app

import (
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
	"github.com/rezaAmiri123/test-microservice/api_service/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser    command.CreateUserHandler
	CreateArticle command.CreateArticleHandler
	CreateComment command.CreateCommentHandler
}

type Queries struct {
	GetArticleBySlug query.GetArticleBySlugHandler
	GetArticles      query.GetArticlesHandler
	GetEmailByUUID   query.GetEmailByUUIDHandler
}
