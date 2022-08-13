package queries

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
)

type GetEmailsHandler struct {
	repo email.Repository
}

func NewGetEmailsHandler(repo email.Repository) *GetEmailsHandler {
	if repo == nil {
		panic("email repo is nil")
	}
	return &GetEmailsHandler{repo: repo}
}

func (h *GetEmailsHandler) Handle(ctx context.Context, page *pagnation.Pagination) (*email.EmailList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetEmailsHandler.Handle")
	defer span.Finish()

	a, err := h.repo.List(ctx, page)
	if err != nil {
		return &email.EmailList{}, err
	}
	return a, nil
}
