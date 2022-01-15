package queries

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
)

type GetArticleBySlugHandler struct {
	repo article.Repository
}

func NewGetArticleHandler(repo article.Repository) *GetArticleBySlugHandler {
	if repo == nil {
		panic("article repo is nil")
	}
	return &GetArticleBySlugHandler{repo: repo}
}

func (h *GetArticleBySlugHandler) Handle(ctx context.Context, slug string) (*article.Article, error) {
	a, err := h.repo.GetBySlug(ctx, slug)
	if err != nil {
		return &article.Article{}, err
	}
	return a, nil
}
