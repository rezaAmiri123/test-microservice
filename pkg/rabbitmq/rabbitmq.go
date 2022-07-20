package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
}

// Initialize new RabbitMQ connection
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
