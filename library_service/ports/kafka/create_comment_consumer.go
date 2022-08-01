package kafka

import (
	"context"

	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	kafkaMessages "github.com/rezaAmiri123/test-microservice/library_service/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

//var (
//	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
//)

func (s *libraryMessageProcessor) processCreateComment(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metric.CreateCommentKafkaRequests.Inc()
	//ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "productMessageProcessor.processCreateProduct")
	//defer span.Finish()

	var msg kafkaMessages.CommentCreateRequest
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	commentObj := &article.Comment{
		UserUUID:    msg.GetUserUUID(),
		UUID:        msg.GetUUID(),
		ArticleUUID: msg.GetArticleUUID(),
		Message:     msg.GetMessage(),
	}
	err := s.app.Commands.CreateComment.Handle(ctx, commentObj, msg.GetUserUUID())
	if err != nil {
		s.log.Errorf("error create article consumer", err)
	}
	s.commitMessage(ctx, r, m)
}
