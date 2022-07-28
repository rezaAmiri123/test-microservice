//go:generate mockgen -source rabbitmq.go -destination mock/rabbitmq.go -package mock
package rabbitmq

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
)

type Config struct {
	User           string
	Password       string
	Host           string
	Port           int
	Exchange       string
	Queue          string
	RoutingKey     string
	ConsumerTag    string
	WorkerPoolSize int
}
type Worker func(ctx context.Context, messages <-chan amqp.Delivery)

// Publisher interface
type Publisher interface {
	Publish(ctx context.Context, body []byte, contentType string) error
}

// Consumer interface
type Consumer interface {
	StartConsumer(ctx context.Context, workerPoolSize int, exchange, queueName, bindingKey, consumerTag string, worker Worker) error
}

// NewRabbitMQConn Initialize new RabbitMQ connection
func NewRabbitMQConn(cfg Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	return amqp.Dial(connAddr)
}
