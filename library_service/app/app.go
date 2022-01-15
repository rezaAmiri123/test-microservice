package app

import (
	"github.com/rezaAmiri123/test-microservice/library_service/app/commands"
	"github.com/rezaAmiri123/test-microservice/library_service/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
	GetArticleBySlug *queries.GetArticleBySlugHandler
}

type Commands struct {
	CreateArticle *commands.CreateArticleHandler
}
