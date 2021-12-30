package query

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
)

type GetUserTokenHandler struct {
	repo domain.Repository
}

func NewGetUserTokenHandler(userRepo domain.Repository) GetUserTokenHandler {
	if userRepo == nil {
		panic("nil userRepo")
	}
	return GetUserTokenHandler{repo: userRepo}
}

func (h GetUserTokenHandler) Handler(ctx context.Context, token string) (*domain.User, error) {
	username, err := domain.GetUsernameFromJWTToken(token)
	if err != nil{
		return &domain.User{}, err
	}
	u,err := h.repo.GetByUsername(ctx,username)
	if err != nil{
		return &domain.User{}, err
	}
	u.HidePassword()
	return u,nil
}
