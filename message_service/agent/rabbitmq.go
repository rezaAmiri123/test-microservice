package agent

import (
	"context"
	messagerabbitmq "github.com/rezaAmiri123/test-microservice/message_service/ports/rabbitmq"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq"
)

func (a *Agent) setupRabbitMQ() error {
	conn, err := rabbitmq.NewRabbitMQConn(a.RabbitmqConfig)
	if err != nil {
		return err
	}
	consumer := messagerabbitmq.NewMessageConsumer(conn, a.logger, a.Application)
	worker := consumer.CreateEmailWorker()
	err = consumer.StartConsumer(
		context.Background(),
		a.RabbitmqConfig.WorkerPoolSize,
		a.RabbitmqConfig.Exchange,
		a.RabbitmqConfig.Queue,
		a.RabbitmqConfig.RoutingKey,
		a.RabbitmqConfig.ConsumerTag,
		worker,
	)
	return err
}
