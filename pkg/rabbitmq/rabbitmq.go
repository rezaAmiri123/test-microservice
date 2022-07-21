//go:generate mockgen -source rabbitmq.go -destination mock/rabbitmq.go -package mock
package rabbitmq

import (
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

// Publisher interface
type Publisher interface {
	Publish(body []byte, contentType string) error
}

// Consumer interface
type Consumer interface {
	StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string) error
}


// NewRabbitMQConn Initialize new RabbitMQ connection
func NewRabbitMQConn(cfg *Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	return amqp.Dial(connAddr)
}
