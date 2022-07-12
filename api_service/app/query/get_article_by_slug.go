package query

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/domain/dto"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
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

	r := &libraryapi.GetArticleRequest{
		Slug: req.Slug,
	}
	a, err := h.client.GetArticle(ctx, r)
	if err != nil {
		return &dto.GetArticleBySlugResponse{}, err
	}

	res := &dto.GetArticleBySlugResponse{
		Title:       a.GetArticle().GetTitle(),
		Body:        a.GetArticle().GetBody(),
		Description: a.GetArticle().GetDescription(),
	}
	return res, nil
}
