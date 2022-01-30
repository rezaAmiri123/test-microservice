package kafka

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	kafkaMessages "github.com/rezaAmiri123/test-microservice/message_service/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

//var (
//	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
//)

func (s *messageMessageProcessor) processCreateEmail(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	//s.metric.CreateUserKafkaRequests.Inc()
	//ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "productMessageProcessor.processCreateProduct")
	//defer span.Finish()

	var msg kafkaMessages.CreateEmail
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	protoEmail := msg.GetEmail()
	e := &email.Email{
		From:    protoEmail.GetFrom(),
		To:      protoEmail.GetTo(),
		Subject: protoEmail.GetSubject(),
		Body:    protoEmail.GetBody(),
	}
	// send email

	err := s.app.Commands.CreateEmail.Handle(ctx, e)
	if err != nil {
		s.log.Errorf("error create user consumer", err)
	}
	s.commitMessage(ctx, r, m)
}
