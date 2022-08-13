package agent

import (
	"fmt"

	"github.com/rezaAmiri123/test-microservice/api_service/app"
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
	"github.com/rezaAmiri123/test-microservice/api_service/app/query"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	messageapi "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
)

func (a *Agent) setupApplication() error {
	producer := kafkaClient.NewProducer(a.logger, a.KafkaConfig.Brokers)

	addr := fmt.Sprintf("%s:%d", a.GRPCLibraryClientAddr, a.GRPCLibraryClientPort)
	libraryConn, err := a.getGrpcClient(addr, a.GRPCLibraryClientTLSConfig)
	if err != nil {
		return err
	}
	articleClient := libraryapi.NewArticleServiceClient(libraryConn)

	addr = fmt.Sprintf("%s:%d", a.GRPCMessageClientAddr, a.GRPCMessageClientPort)
	messageConn, err := a.getGrpcClient(addr, a.GRPCMessageClientTLSConfig)
	if err != nil {
		return err
	}
	messageClient := messageapi.NewMessageServiceClient(messageConn)

	application := &app.Application{
		Commands: app.Commands{
			CreateUser:    command.NewCreateUserHandler(producer, a.logger),
			CreateArticle: command.NewCreateArticleHandler(producer, a.logger),
			CreateComment: command.NewCreateCommentHandler(producer, a.logger),
		},
		Queries: app.Queries{
			GetArticleBySlug: query.NewGetArticleBySlugHandler(articleClient, a.logger),
			GetArticles:      query.NewGetArticlesHandler(articleClient, a.logger),
			GetEmailByUUID:   query.NewGetEmailByUUIDHandler(messageClient, a.logger),
			GetEmails:        query.NewGetEmailsHandler(messageClient, a.logger),
		},
	}
	a.Application = application
	return nil
}
