package kafka

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/metrics"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/segmentio/kafka-go"
	"sync"
)

const PoolSize = 30

type libraryMessageProcessor struct {
	log    logger.Logger
	cfg    Config
	metric *metrics.LibraryServiceMetric
	app    *app.Application
}

func NewLibraryMessageProcessor(log logger.Logger, cfg Config, metric *metrics.LibraryServiceMetric, app *app.Application) *libraryMessageProcessor {
	return &libraryMessageProcessor{log: log, cfg: cfg, metric: metric, app: app}
}

func (s *libraryMessageProcessor) ProcessMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
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
		case s.cfg.KafkaTopics.ArticleCreate.TopicName:
			s.processCreateArticle(ctx, r, m)
		case s.cfg.KafkaTopics.CommentCreate.TopicName:
			s.processCreateComment(ctx, r, m)
		}
	}
}
