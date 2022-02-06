package command

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	kafkaMessages "github.com/rezaAmiri123/test-microservice/message_service/proto/kafka"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type CreateUserHandler struct {
	logger   logger.Logger
	userRepo domain.Repository
	kafka    kafkaClient.Producer
}

func NewCreateUserHandler(userRepo domain.Repository, kafka kafkaClient.Producer, logger logger.Logger) CreateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return CreateUserHandler{userRepo: userRepo, kafka: kafka, logger: logger}
}

func (h CreateUserHandler) Handle(ctx context.Context, user *domain.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.Handle")
	defer span.Finish()

	if err := user.SetUUID(); err != nil {
		return err
	}
	if err := user.Validate(ctx); err != nil {
		return err
	}
	if err := user.HashPassword(); err != nil {
		return err
	}

	e := &kafkaMessages.Email{
		To:      []string{user.Email},
		From:    "admin@example.com",
		Subject: "register user subject",
		Body:    "register user body",
	}
	msg := &kafkaMessages.CreateEmail{Email: e}

	message, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	err = h.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	err = h.kafka.PublishMessage(ctx, kafka.Message{
		Topic: kafkaClient.CreateEmailTopic,
		Value: message,
		Time:  time.Now().UTC(),
	})
	if err != nil {
		h.logger.Errorf("can not send kafka message %v", err)
	}
	return err
}
