package queries

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
)

type GetEmailBySlugHandler struct {
	Repo email.Repository
}

func NewGetArticleHandler(repo email.Repository) *GetEmailBySlugHandler {
	if repo == nil {
		panic("article repo is nil")
	}
	return &GetEmailBySlugHandler{Repo: repo}
}

func (h *GetEmailBySlugHandler) Handle(ctx context.Context, slug string) (*email.Email, error) {
	e, err := h.Repo.GetByUUID(ctx, slug)
	if err != nil {
		return &email.Email{}, err
	}
	return e, nil
}
