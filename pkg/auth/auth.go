package auth

import (
	"context"

	UserApi "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
)

type AuthClient interface {
	Login(ctx context.Context, username, password string) (string, error)
	VerityToken(ctx context.Context, token string) (*User, error)
}

type User struct {
	Username string
	UUID     string
}

func NewUserAuthClient(authClient UserApi.UsersServiceClient) (AuthClient, error) {
	return &UserAuthClient{AuthClient: authClient}, nil
}

type UserAuthClient struct {
	AuthClient UserApi.UsersServiceClient
}

func (a *UserAuthClient) Login(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (a *UserAuthClient) VerityToken(ctx context.Context, token string) (*User, error) {
	return nil, nil
}
