package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
)

type CreateEmailHandler struct {
	Repo email.Repository
}

func NewCreateUserHandler(repo email.Repository) *CreateEmailHandler {
	if repo == nil {
		panic("email Repo is nil")
	}
	return &CreateEmailHandler{Repo: repo}
}

func (h CreateEmailHandler) Handle(ctx context.Context, e *email.Email) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateEmailHandler.Handle")
	defer span.Finish()

	if err := e.Validate(ctx); err != nil {
		return err
	}

	err := h.Repo.Create(ctx, e)
	return err
}
