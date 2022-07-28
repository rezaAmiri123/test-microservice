package commands

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	"github.com/rezaAmiri123/test-microservice/pkg/converter"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq"
)

type CreateEmailWithQueueHandler struct {
	Publisher rabbitmq.Publisher
}

func NewCreateEmailWithQueueHandler(publisher rabbitmq.Publisher) *CreateEmailWithQueueHandler {
	if publisher == nil {
		panic("email Publisher is nil")
	}
	return &CreateEmailWithQueueHandler{Publisher: publisher}
}

func (h CreateEmailWithQueueHandler) Handle(ctx context.Context, e *email.Email) (*email.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateEmailWithQueueHandler.Handle")
	defer span.Finish()
	if err := e.SetUUID(); err != nil {
		return nil, err
	}

	if err := e.Validate(ctx); err != nil {
		return nil, err
	}
	buf, err := converter.AnyToBytesBuffer(e)
	if err != nil {
		return nil, err
	}

	contentType := "application/json"
	err = h.Publisher.Publish(ctx, buf.Bytes(), contentType)
	return e, err
}
