package rabbitmq

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	"github.com/rezaAmiri123/test-microservice/pkg/converter"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func (c *MessageConsumer) CreateEmailWorker() rabbitmq.Worker {
	return func(ctx context.Context, messages <-chan amqp.Delivery) {
		for delivery := range messages {
			span, ctx := opentracing.StartSpanFromContext(ctx, "EmailsConsumer.worker")

			c.logger.Infof("processDeliveries deliveryTag% v", delivery.DeliveryTag)

			//incomingMessages.Inc()
			emailByte, err := converter.BytesToStruct(delivery.Body)
			e := emailByte.(email.Email)
			err = c.app.Commands.CreateEmail.Handle(ctx, &e)
			//if err != nil {
			//	s.log.Errorf("error create user consumer", err)
			//}
			//err := c.emailUC.SendEmail(ctx, delivery.Body)
			if err != nil {
				if err := delivery.Reject(false); err != nil {
					c.logger.Errorf("Err delivery.Reject: %v", err)
				}
				c.logger.Errorf("Failed to process delivery: %v", err)
				//errorMessages.Inc()
			} else {
				err = delivery.Ack(false)
				if err != nil {
					c.logger.Errorf("Failed to acknowledge delivery: %v", err)
				}
				//successMessages.Inc()
			}
			span.Finish()
		}

		c.logger.Info("Deliveries channel closed")
	}
}
