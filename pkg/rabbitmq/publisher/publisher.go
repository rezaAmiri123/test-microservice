package publisher

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

const (
	exchangeKind       = "direct"
	exchangeDurable    = true
	exchangeAutoDelete = false
	exchangeInternal   = false
	exchangeNoWait     = false

	queueDurable    = true
	queueAutoDelete = false
	queueExclusive  = false
	queueNoWait     = false

	publishMandatory = false
	publishImmediate = false

	prefetchCount  = 1
	prefetchSize   = 0
	prefetchGlobal = false

	consumeAutoAck   = false
	consumeExclusive = false
	consumeNoLocal   = false
	consumeNoWait    = false
)

//var (
//	publishedMessages = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "emails_published_rabbitmq_messages_total",
//		Help: "The total number of published RabbitMQ messages",
//	})
//)

// Publisher rabbitmq publisher
type Publisher struct {
	amqpChan *amqp.Channel
	cfg      *rabbitmq.Config
	logger   logger.Logger
}

// NewEmailsPublisher Emails rabbitmq publisher constructor
func NewPublisher(cfg *rabbitmq.Config, logger logger.Logger) (*Publisher, error) {
	mqConn, err := rabbitmq.NewRabbitMQConn(cfg)
	if err != nil {
		return nil, err
	}
	amqpChan, err := mqConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "p.amqpConn.Channel")
	}

	return &Publisher{cfg: cfg, logger: logger, amqpChan: amqpChan}, nil
}

// SetupExchangeAndQueue create exchange and queue
func (p *Publisher) SetupExchangeAndQueue(exchange, queueName, bindingKey, consumerTag string) error {
	p.logger.Infof("Declaring exchange: %s", exchange)
	err := p.amqpChan.ExchangeDeclare(
		exchange,
		exchangeKind,
		exchangeDurable,
		exchangeAutoDelete,
		exchangeInternal,
		exchangeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.ExchangeDeclare")
	}

	queue, err := p.amqpChan.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		queueNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.QueueDeclare")
	}

	p.logger.Infof("Declared queue, binding it to exchange: Queue: %v, messageCount: %v, "+
		"consumerCount: %v, exchange: %v, exchange: %v, bindingKey: %v",
		queue.Name,
		queue.Messages,
		queue.Consumers,
		exchange,
		bindingKey,
	)

	err = p.amqpChan.QueueBind(
		queue.Name,
		bindingKey,
		exchange,
		queueNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.QueueBind")
	}

	p.logger.Infof("Queue bound to exchange, starting to consume from queue, consumerTag: %v", consumerTag)
	return nil
}

// CloseChan Close messages chan
func (p *Publisher) CloseChan() {
	if err := p.amqpChan.Close(); err != nil {
		p.logger.Errorf("EmailsPublisher CloseChan: %v", err)
	}
}

// Publish message
func (p *Publisher) Publish(body []byte, contentType string) error {

	p.logger.Infof("Publishing message Exchange: %s, RoutingKey: %s", p.cfg.Exchange, p.cfg.RoutingKey)

	if err := p.amqpChan.Publish(
		p.cfg.Exchange,
		p.cfg.RoutingKey,
		publishMandatory,
		publishImmediate,
		amqp.Publishing{
			ContentType:  contentType,
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Body:         body,
		},
	); err != nil {
		return errors.Wrap(err, "ch.Publish")
	}

	//publishedMessages.Inc()
	return nil
}
