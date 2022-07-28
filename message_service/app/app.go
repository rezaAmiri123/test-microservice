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
	GetEmailByUUID *queries.GetEmailByUUIDHandler
}

type Commands struct {
	CreateEmail          *commands.CreateEmailHandler
	CreateEmailWithQueue *commands.CreateEmailWithQueueHandler
}
