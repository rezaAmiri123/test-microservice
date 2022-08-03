package command

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	kafkaLibrary "github.com/rezaAmiri123/test-microservice/library_service/proto/kafka"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type CreateCommentHandler struct {
	logger logger.Logger
	kafka  kafkaClient.Producer
}

func NewCreateCommentHandler(kafka kafkaClient.Producer, logger logger.Logger) CreateCommentHandler {
	return CreateCommentHandler{kafka: kafka, logger: logger}
}

func (h CreateCommentHandler) Handle(ctx context.Context, comment *dto.CreateCommentRequest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateCommentHandler.Handle")
	defer span.Finish()

	commentObj := &kafkaLibrary.CommentCreateRequest{
		UUID:        comment.UUID,
		UserUUID:    comment.UserUUID,
		ArticleUUID: comment.ArticleUUID,
		Message:     comment.Message,
	}

	message, err := proto.Marshal(commentObj)
	if err != nil {
		return err
	}
	err = h.kafka.PublishMessage(ctx, kafka.Message{
		Topic:   kafkaClient.CreateCommentTopic,
		Value:   message,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
	if err != nil {
		h.logger.Errorf("can not send kafka message %v", err)
	}
	return err
}
