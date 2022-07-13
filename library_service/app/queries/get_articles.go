package queries

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
)

type GetArticlesHandler struct {
	repo article.Repository
}

func NewGetArticlesHandler(repo article.Repository) *GetArticlesHandler {
	if repo == nil {
		panic("article repo is nil")
	}
	return &GetArticlesHandler{repo: repo}
}

func (h *GetArticlesHandler) Handle(ctx context.Context, page *pagnation.Pagination) (*article.ArticleList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetArticlesHandler.Handle")
	defer span.Finish()

	a, err := h.repo.List(ctx, page)
	if err != nil {
		return &article.ArticleList{}, err
	}
	return a, nil
}
