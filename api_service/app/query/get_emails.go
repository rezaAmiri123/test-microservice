package query

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	messageapi "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type GetEmailsHandler struct {
	client messageapi.MessageServiceClient
	logger logger.Logger
}

func NewGetEmailsHandler(
	client messageapi.MessageServiceClient,
	logger logger.Logger,
) GetEmailsHandler {
	if client == nil {
		panic("article client is nil")
	}
	return GetEmailsHandler{client: client, logger: logger}
}

func (h *GetEmailsHandler) Handle(ctx context.Context, query *pagnation.Pagination) (*dto.GetEmailsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetArticlesHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	req := &messageapi.GetEmailsRequest{
		Page: int64(query.GetPage()),
		Size: int64(query.GetSize()),
	}
	list, err := h.client.GetEmails(ctx, req)
	if err != nil {
		return &dto.GetEmailsResponse{}, err
	}

	res := &dto.GetEmailsResponse{}
	res.TotalCount = list.GetTotalCount()
	res.TotalPages = list.GetTotalPages()
	res.Page = list.GetPage()
	res.Size = list.GetSize()
	res.HasMore = list.GetHasMore()

	emailList := make([]*dto.EmailResponse, 0, len(list.GetEmails()))
	for _, a := range list.GetEmails() {
		emailList = append(emailList, dto.EmailResponseFromGrpc(a))
	}
	res.Emails = emailList

	return res, nil
}
