package queries

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
)

type GetEmailByUUIDHandler struct {
	Repo email.Repository
}

func NewGetEmailHandler(repo email.Repository) *GetEmailByUUIDHandler {
	if repo == nil {
		panic("article repo is nil")
	}
	return &GetEmailByUUIDHandler{Repo: repo}
}

func (h *GetEmailByUUIDHandler) Handle(ctx context.Context, slug string) (*email.Email, error) {
	e, err := h.Repo.GetByUUID(ctx, slug)
	if err != nil {
		return &email.Email{}, err
	}
	return e, nil
}
