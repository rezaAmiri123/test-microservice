package query

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type GetArticlesHandler struct {
	client libraryapi.ArticleServiceClient
	logger logger.Logger
}

func NewGetArticlesHandler(
	client libraryapi.ArticleServiceClient,
	logger logger.Logger,
) GetArticlesHandler {
	if client == nil {
		panic("article client is nil")
	}
	return GetArticlesHandler{client: client, logger: logger}
}

func (h *GetArticlesHandler) Handle(ctx context.Context, query *pagnation.Pagination) (*dto.GetArticlesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetArticlesHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	req := &libraryapi.GetArticlesRequest{
		Page: int64(query.GetPage()),
		Size: int64(query.GetSize()),
	}
	list, err := h.client.GetArticles(ctx, req)
	if err != nil {
		return &dto.GetArticlesResponse{}, err
	}

	res := &dto.GetArticlesResponse{}
	res.TotalCount = list.GetTotalCount()
	res.TotalPages = list.GetTotalPages()
	res.Page = list.GetPage()
	res.Size = list.GetSize()
	res.HasMore = list.GetHasMore()

	articleList := make([]*dto.ArticleResponse, 0, len(list.GetArticles()))
	for _, a := range list.GetArticles() {
		articleList = append(articleList, dto.ArticleResponseFromGrpc(a))
	}
	res.Articles = articleList

	return res, nil
}
