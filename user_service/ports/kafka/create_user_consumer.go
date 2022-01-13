package kafka

import (
	"context"
	"github.com/avast/retry-go"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	kafkaMessages "github.com/rezaAmiri123/test-microservice/user_service/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (s *userMessageProcessor) processCreateUser(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metric.CreateUserKafkaRequests.Inc()
	//ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "productMessageProcessor.processCreateProduct")
	//defer span.Finish()

	var msg kafkaMessages.CreateUser
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	user := msg.GetUser()
	//if err := retry.Do(func() error {
	//	//return s.ps.Commands.CreateProduct.Handle(ctx, command)
	//	return s.app.Commands.CreateUser.Handle(ctx, &domain.User{
	//		Username: user.GetUsername(),
	//		Password: user.GetPassword(),
	//		Email:    user.GetEmail(),
	//		Bio:      user.GetBio(),
	//		Image:    user.GetImage(),
	//	})
	//}, append(retryOptions, retry.Context(ctx))...); err != nil {
	//	s.log.WarnMsg("CreateUser.Handle", err)
	//	s.metric.ErrorKafkaMessages.Inc()
	//	s.log.Info("user not created llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll")
	//	return
	//}
	err := s.app.Commands.CreateUser.Handle(ctx, &domain.User{
		Username: user.GetUsername(),
		Password: user.GetPassword(),
		Email:    user.GetEmail(),
		Bio:      user.GetBio(),
		Image:    user.GetImage(),
	})
	if err != nil{
		s.log.Errorf("error create user consumer", err)
	}
	s.log.Info("user created llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll")
	s.commitMessage(ctx, r, m)

}