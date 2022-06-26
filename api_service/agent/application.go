package agent

import (
	"github.com/rezaAmiri123/test-microservice/api_service/app"
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
)

func (a *Agent) setupApplication() error {
	//dbConn, err := postgres.NewPsqlDB(a.DBConfig)
	//if err != nil {
	//	return err
	//}
	//
	////repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	//repo := pg.NewPGArticleRepository(dbConn)
	producer := kafkaClient.NewProducer(a.logger, a.KafkaConfig.Brokers)
	application := &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(producer, a.logger),
		},
		//Queries: app.Queries{
		//	GetArticleBySlug: queries.NewGetArticleHandler(repo),
		//},
	}
	a.Application = application
	return nil
}
