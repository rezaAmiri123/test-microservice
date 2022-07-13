package command

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	kafkaUser "github.com/rezaAmiri123/test-microservice/user_service/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type CreateUserHandler struct {
	logger logger.Logger
	kafka  kafkaClient.Producer
}

func NewCreateUserHandler(kafka kafkaClient.Producer, logger logger.Logger) CreateUserHandler {
	return CreateUserHandler{kafka: kafka, logger: logger}
}

func (h CreateUserHandler) Handle(ctx context.Context, user *dto.CreateUserRequest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.Handle")
	defer span.Finish()

	u := &kafkaUser.User{
		UUID:     user.UserID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Bio:      user.Bio,
		Image:    user.Image,
	}
	msg := &kafkaUser.CreateUser{User: u}
	message, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	err = h.kafka.PublishMessage(ctx, kafka.Message{
		Topic:   kafkaClient.CreateUserTopic,
		Value:   message,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
	if err != nil {
		h.logger.Errorf("can not send kafka message %v", err)
	}
	return err
}
