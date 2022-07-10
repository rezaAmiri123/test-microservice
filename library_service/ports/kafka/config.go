package kafka

import (
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
)

type Config struct {
	KafkaTopics KafkaTopics
	Kafka       kafkaClient.Config
}
type KafkaTopics struct {
	ArticleCreate kafkaClient.TopicConfig
}
