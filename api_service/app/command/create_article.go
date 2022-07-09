package command

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/domain/dto"
	kafkaLibrary "github.com/rezaAmiri123/test-microservice/library_service/proto/kafka"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type CreateArticleHandler struct {
	logger logger.Logger
	kafka  kafkaClient.Producer
}

func NewCreateArticleHandler(kafka kafkaClient.Producer, logger logger.Logger) CreateArticleHandler {
	return CreateArticleHandler{kafka: kafka, logger: logger}
}

func (h CreateArticleHandler) Handle(ctx context.Context, article *dto.CreateArticleRequest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateArticleHandler.Handle")
	defer span.Finish()

	articleObj := &kafkaLibrary.ArticleCreateRequest{
		UUID:        article.ArticleID,
		UserUUID:    article.UserID,
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
	}

	message, err := proto.Marshal(articleObj)
	if err != nil {
		return err
	}
	err = h.kafka.PublishMessage(ctx, kafka.Message{
		Topic:   kafkaClient.CreateArticleTopic,
		Value:   message,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
	if err != nil {
		h.logger.Errorf("can not send kafka message %v", err)
	}
	return err
}
