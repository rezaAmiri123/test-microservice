package app

import (
	"github.com/rezaAmiri123/test-microservice/message_service/app/commands"
	"github.com/rezaAmiri123/test-microservice/message_service/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
	GetArticleBySlug *queries.GetEmailBySlugHandler
}

type Commands struct {
	CreateArticle *commands.CreateEmailHandler
}
