package agent

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/message_service/ports/kafka"
	kafkaClient "github.com/rezaAmiri123/test-microservice/pkg/kafka"
)

func (a *Agent) setupKafka() error {
	ctx, cancel := context.WithCancel(context.Background())
	a.closers = append(a.closers, closer{cancel: cancel})
	messageMessageProcessor := kafka.NewMessageMessageProcessor(a.logger, a.KafkaConfig, a.metric, a.Application)
	cg := kafkaClient.NewConsumerGroup(a.KafkaConfig.Kafka.Brokers, a.KafkaConfig.Kafka.GroupID, a.logger)
	//kafkaConn, err := kafkaClient.NewKafkaConn(ctx, &a.KafkaConfig.Kafka)
	//if err != nil {
	//	return errors.Wrap(err, "kafka.NewKafkaCon")
	//}
	//
	topics := []string{
		a.KafkaConfig.KafkaTopics.EmailCreate.TopicName,
	}
	// TODO we need a context to can ended goroutine
	go cg.ConsumeTopic(ctx, topics, kafka.PoolSize, messageMessageProcessor.ProcessMessage)

	return nil
}

type closer struct {
	cancel func()
}

func (c closer) Close() error {
	c.cancel()
	return nil
}
