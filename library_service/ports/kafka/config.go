package kafka

import (
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

type Config struct {
	KafkaTopics KafkaTopics
	Kafka       kafkaClient.Config
}
type KafkaTopics struct {
	ArticleCreate kafkaClient.TopicConfig
	CommentCreate kafkaClient.TopicConfig
}
