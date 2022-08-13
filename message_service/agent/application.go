package agent

import (
	"context"

	"github.com/rezaAmiri123/test-microservice/message_service/adapters/mongo"
	"github.com/rezaAmiri123/test-microservice/message_service/app"
	"github.com/rezaAmiri123/test-microservice/message_service/app/commands"
	"github.com/rezaAmiri123/test-microservice/message_service/app/queries"
	"github.com/rezaAmiri123/test-microservice/pkg/mongodb"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq/publisher"
)

func (a *Agent) setupApplication() error {
	//dbConn, err := postgres.NewPsqlDB(a.DBConfig)
	//if err != nil {
	//	return err
	//}

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	//repo := pg.NewPGEmailRepository(dbConn)

	dbConn, err := mongodb.NewMongoDBConn(context.Background(), &a.MongoConfig)
	if err != nil {
		return err
	}

	repo := mongo.NewMongoRepository(a.logger, &a.MongoConfig, dbConn)
	rabbitPublisher, err := publisher.NewPublisher(a.RabbitmqConfig, a.logger)
	application := &app.Application{
		Commands: app.Commands{
			CreateEmail:          commands.NewCreateEmailHandler(repo),
			CreateEmailWithQueue: commands.NewCreateEmailWithQueueHandler(rabbitPublisher),
		},
		Queries: app.Queries{
			GetEmailByUUID: queries.NewGetEmailHandler(repo),
			GetEmails:      queries.NewGetEmailsHandler(repo),
		},
	}
	a.Application = application
	return nil
}
