package kafka

import (
	"context"

	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	kafkaMessages "github.com/rezaAmiri123/test-microservice/library_service/proto/kafka"
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

func (s *libraryMessageProcessor) processCreateArticle(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metric.CreateArticleKafkaRequests.Inc()
	//ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "productMessageProcessor.processCreateProduct")
	//defer span.Finish()

	var msg kafkaMessages.ArticleCreateRequest
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	articleObj := &article.Article{
		UserUUID:    msg.GetUserUUID(),
		UUID:        msg.GetUUID(),
		Title:       msg.GetTitle(),
		Description: msg.GetDescription(),
		Body:        msg.GetBody(),
	}
	err := s.app.Commands.CreateArticle.Handle(ctx, articleObj, msg.GetUserUUID())
	if err != nil {
		s.log.Errorf("error create article consumer", err)
	}
	s.commitMessage(ctx, r, m)
}
