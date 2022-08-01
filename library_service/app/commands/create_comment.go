package commands

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
)

type CreateCommentHandler struct {
	repo article.Repository
}

func NewCreateCommentHandler(repo article.Repository) *CreateCommentHandler {
	if repo == nil {
		panic("repo is nil")
	}
	return &CreateCommentHandler{repo: repo}
}

func (h CreateCommentHandler) Handle(ctx context.Context, c *article.Comment, userUUID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateArticleHandler.Handle")
	defer span.Finish()

	c.UserUUID = userUUID
	c.Likes = 0
	if err := c.Validate(ctx); err != nil {
		return err
	}

	err := h.repo.CreateComment(ctx, c)
	return err
}
