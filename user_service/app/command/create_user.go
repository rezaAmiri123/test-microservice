package command

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
)

type CreateUserHandler struct {
	userRepo domain.Repository
}

func NewCreateUserHandler(userRepo domain.Repository) CreateUserHandler {
	if userRepo==nil{
		panic("userRepo is nil")
	}
	return CreateUserHandler{userRepo: userRepo}
}

func (h CreateUserHandler) Handle(ctx context.Context, user *domain.User)error {
	if err := user.SetUUID(); err != nil {
		return err
	}
	if err := user.Validate(); err != nil {
		return err
	}
	if err := user.HashPassword(); err != nil {
		return err
	}
	err := h.userRepo.Create(ctx, user)
	return err
}