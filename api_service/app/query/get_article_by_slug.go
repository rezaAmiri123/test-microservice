package query

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type GetArticleBySlugHandler struct {
	client libraryapi.ArticleServiceClient
	logger logger.Logger
}

func NewGetArticleBySlugHandler(
	client libraryapi.ArticleServiceClient,
	logger logger.Logger,
) GetArticleBySlugHandler {
	if client == nil {
		panic("article client is nil")
	}
	return GetArticleBySlugHandler{client: client, logger: logger}
}

func (h *GetArticleBySlugHandler) Handle(ctx context.Context, req *dto.GetArticleBySlugRequest) (*dto.GetArticleBySlugResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetArticleBySlugHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	r := &libraryapi.GetArticleBySlugRequest{
		Slug: req.Slug,
	}
	a, err := h.client.GetArticleBySlug(ctx, r)
	if err != nil {
		return &dto.GetArticleBySlugResponse{}, err
	}

	res := &dto.GetArticleBySlugResponse{}
	res.Title = a.GetArticle().GetTitle()
	res.Body = a.GetArticle().GetBody()
	res.Description = a.GetArticle().GetDescription()

	return res, nil
}
