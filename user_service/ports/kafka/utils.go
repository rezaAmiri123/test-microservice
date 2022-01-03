package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func (s *userMessageProcessor) logProcessMessage(m kafka.Message, workerID int) {
	s.log.KafkaProcessMessage(m.Topic,m.Partition, string(m.Value), workerID, m.Offset, m.Time)
}

func (s *userMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metric.SuccessKafkaMessages.Inc()
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition,m.Offset)
	if err := r.CommitMessages(ctx,m);err!= nil{
		s.log.WarnMsg("commitMessage", err)
	}
}

func (s *userMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metric.ErrorKafkaMessages.Inc()
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition,m.Offset)
	if err := r.CommitMessages(ctx,m);err!= nil{
		s.log.WarnMsg("commitMessage", err)
	}
}