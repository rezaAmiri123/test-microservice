package query

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
)

type GetProfileHandler struct {
	userRepo domain.Repository
}

func NewGetProfileHandler(userRepo domain.Repository) GetProfileHandler {
	if userRepo == nil {
		panic("nil userRepo")
	}
	return GetProfileHandler{userRepo: userRepo}
}

func (h GetProfileHandler) Handle(ctx context.Context, username string) (*domain.User, error) {
	u, err := h.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return &domain.User{}, err
	}
	u.HidePassword()
	return u, nil
}
