package agent

import (
	"github.com/rezaAmiri123/test-microservice/api_service/app"
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
	"github.com/rezaAmiri123/test-microservice/api_service/app/query"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
)

func (a *Agent) setupApplication() error {
	producer := kafkaClient.NewProducer(a.logger, a.KafkaConfig.Brokers)
	libraryConn, err := a.getLibraryClient()
	if err != nil {
		return err
	}
	articleClient := libraryapi.NewArticleServiceClient(libraryConn)

	application := &app.Application{
		Commands: app.Commands{
			CreateUser:    command.NewCreateUserHandler(producer, a.logger),
			CreateArticle: command.NewCreateArticleHandler(producer, a.logger),
		},
		Queries: app.Queries{
			GetArticleBySlug: query.NewGetArticleBySlugHandler(articleClient, a.logger),
		},
	}
	a.Application = application
	return nil
}
