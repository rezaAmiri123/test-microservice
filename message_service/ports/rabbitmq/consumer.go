package rabbitmq

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/message_service/app"
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
//	incomingMessages = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "emails_incoming_rabbitmq_messages_total",
//		Help: "The total number of incoming RabbitMQ messages",
//	})
//	successMessages = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "emails_success_incoming_rabbitmq_messages_total",
//		Help: "The total number of success incoming success RabbitMQ messages",
//	})
//	errorMessages = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "emails_error_incoming_rabbitmq_message_total",
//		Help: "The total number of error incoming success RabbitMQ messages",
//	})
//)

// MessageConsumer Images Rabbitmq consumer
type MessageConsumer struct {
	amqpConn *amqp.Connection
	logger   logger.Logger
	app      *app.Application
}

// NewMessageConsumer Images Consumer constructor
func NewMessageConsumer(amqpConn *amqp.Connection, logger logger.Logger, app *app.Application) *MessageConsumer {
	return &MessageConsumer{amqpConn: amqpConn, logger: logger, app: app}
}

// CreateChannel Consume messages
func (c *MessageConsumer) CreateChannel(exchangeName, queueName, bindingKey, consumerTag string) (*amqp.Channel, error) {
	ch, err := c.amqpConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Error amqpConn.Channel")
	}

	c.logger.Infof("Declaring exchange: %s", exchangeName)
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeKind,
		exchangeDurable,
		exchangeAutoDelete,
		exchangeInternal,
		exchangeNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.ExchangeDeclare")
	}

	queue, err := ch.QueueDeclare(
		queueName,
		queueDurable,
		queueAutoDelete,
		queueExclusive,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueDeclare")
	}

	c.logger.Infof("Declared queue, binding it to exchange: Queue: %v, messagesCount: %v, "+
		"consumerCount: %v, exchange: %v, bindingKey: %v",
		queue.Name,
		queue.Messages,
		queue.Consumers,
		exchangeName,
		bindingKey,
	)

	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchangeName,
		queueNoWait,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueBind")
	}

	c.logger.Infof("Queue bound to exchange, starting to consume from queue, consumerTag: %v", consumerTag)

	err = ch.Qos(
		prefetchCount,  // prefetch count
		prefetchSize,   // prefetch size
		prefetchGlobal, // global
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error  ch.Qos")
	}

	return ch, nil
}

//func (c *MessageConsumer) worker(ctx context.Context, messages <-chan amqp.Delivery) {
//
//	for delivery := range messages {
//		span, ctx := opentracing.StartSpanFromContext(ctx, "EmailsConsumer.worker")
//
//		c.logger.Infof("processDeliveries deliveryTag% v", delivery.DeliveryTag)
//
//		//incomingMessages.Inc()
//
//		err := c.emailUC.SendEmail(ctx, delivery.Body)
//		if err != nil {
//			if err := delivery.Reject(false); err != nil {
//				c.logger.Errorf("Err delivery.Reject: %v", err)
//			}
//			c.logger.Errorf("Failed to process delivery: %v", err)
//			//errorMessages.Inc()
//		} else {
//			err = delivery.Ack(false)
//			if err != nil {
//				c.logger.Errorf("Failed to acknowledge delivery: %v", err)
//			}
//			//successMessages.Inc()
//		}
//		span.Finish()
//	}
//
//	c.logger.Info("Deliveries channel closed")
//}

// StartConsumer Start new rabbitmq consumer
func (c *MessageConsumer) StartConsumer(ctx context.Context, workerPoolSize int, exchange, queueName, bindingKey, consumerTag string, worker rabbitmq.Worker) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch, err := c.CreateChannel(exchange, queueName, bindingKey, consumerTag)
	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}
	defer ch.Close()

	deliveries, err := ch.Consume(
		queueName,
		consumerTag,
		consumeAutoAck,
		consumeExclusive,
		consumeNoLocal,
		consumeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Consume")
	}

	for i := 0; i < workerPoolSize; i++ {
		//go c.worker(ctx, deliveries)
		go worker(ctx, deliveries)
	}

	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	c.logger.Errorf("ch.NotifyClose: %v", chanErr)
	return chanErr
}
