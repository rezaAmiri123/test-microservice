package kafka

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/message_service/app"
	"github.com/rezaAmiri123/test-microservice/message_service/metrics"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/segmentio/kafka-go"
	"sync"
)

const PoolSize = 30

type messageMessageProcessor struct {
	log    logger.Logger
	cfg    Config
	metric *metrics.MessageServiceMetric
	app    *app.Application
}

func NewMessageMessageProcessor(log logger.Logger, cfg Config, metric *metrics.MessageServiceMetric, app *app.Application) *messageMessageProcessor {
	return &messageMessageProcessor{log: log, cfg: cfg, metric: metric, app: app}
}

func (s *messageMessageProcessor) ProcessMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
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
		case s.cfg.KafkaTopics.EmailCreate.TopicName:
			s.processCreateEmail(ctx, r, m)
		}
	}
}
