package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func (s *userMessageProcessor) processCreateUser(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.log.Debug("received user create message %s", string(m.Value))
}