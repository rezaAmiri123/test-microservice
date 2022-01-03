package kafka

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/user_service/metrics"
	"github.com/segmentio/kafka-go"
	"sync"
)

const PoolSize = 30

type userMessageProcessor struct {
	log    logger.Logger
	cfg    Config
	metric *metrics.UserServiceMetric
}

func NewUserMessageProcessor(log logger.Logger, cfg Config, metric *metrics.UserServiceMetric) *userMessageProcessor {
	return &userMessageProcessor{log: log, cfg: cfg, metric: metric}
}

func (s *userMessageProcessor) ProcessMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}
		s.logProcessMessage(m, workerID)
		switch m.Topic {
		case s.cfg.KafkaTopics.UserCreate.TopicName:
			s.processCreateUser(ctx, r, m)
		}
	}
}
