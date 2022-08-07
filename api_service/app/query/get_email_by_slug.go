package query

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	messageapi "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type GetEmailByUUIDHandler struct {
	client messageapi.MessageServiceClient
	logger logger.Logger
}

func NewGetEmailByUUIDHandler(
	client messageapi.MessageServiceClient,
	logger logger.Logger,
) GetEmailByUUIDHandler {
	if client == nil {
		panic("article client is nil")
	}
	return GetEmailByUUIDHandler{client: client, logger: logger}
}

func (h *GetEmailByUUIDHandler) Handle(ctx context.Context, req *dto.GetEmailByUUIDRequest) (*dto.GetEmailByUUIDResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetEmailByUUIDHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	r := &messageapi.GetEmailByUUIDRequest{
		UUID: req.UUID,
	}
	a, err := h.client.GetEmailByUUID(ctx, r)
	if err != nil {
		return &dto.GetEmailByUUIDResponse{}, err
	}

	res := &dto.GetEmailByUUIDResponse{}
	res.Subject = a.GetSubject()
	res.To = a.GetTo()
	res.From = a.GetFrom()
	res.Body = a.GetBody()

	return res, nil
}
