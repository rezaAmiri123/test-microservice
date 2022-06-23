package command

import (
	"context"
	"github.com/opentracing/opentracing-go"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	kafkaUser "github.com/rezaAmiri123/test-microservice/user_service/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type CreateUserHandler struct {
	logger logger.Logger
	kafka  kafkaClient.Producer
}

func (h CreateUserHandler) Handle(ctx context.Context, user *domain.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.Handle")
	defer span.Finish()

	//e := &kafkaMessages.Email{
	//	To:      []string{user.Email},
	//	From:    "admin@example.com",
	//	Subject: "register user subject",
	//	Body:    "register user body",
	//}
	//msg := &kafkaMessages.CreateEmail{Email: e}
	u := &kafkaUser.User{}
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
