package commands

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
)

type CreateArticleHandler struct {
	repo article.Repository
}

func NewCreateArticleHandler(repo article.Repository) *CreateArticleHandler {
	if repo == nil {
		panic("repo is nil")
	}
	return &CreateArticleHandler{repo: repo}
}

func (h CreateArticleHandler) Handle(ctx context.Context, a *article.Article, userUUID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateArticleHandler.Handle")
	defer span.Finish()

	if err := a.SetUUID(userUUID); err != nil {
		return err
	}
	a.Slug = slug.Make(a.Title)
	if err := a.Validate(ctx); err != nil {
		return err
	}

	err := h.repo.Create(ctx, a)
	return err
}
