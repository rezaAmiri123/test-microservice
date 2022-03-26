package auth_client

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	UserApi "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
)

func NewUserAuthClient(authClient UserApi.UsersServiceClient) (auth.AuthClient, error) {
	return &UserAuthClient{AuthClient: authClient}, nil
}

type UserAuthClient struct {
	AuthClient UserApi.UsersServiceClient
}

func (a *UserAuthClient) Login(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (a *UserAuthClient) VerityToken(ctx context.Context, token string) (*auth.User, error) {
	u, err := a.AuthClient.VerifyToken(ctx, &UserApi.VerifyTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}
	return &auth.User{
		UUID:     u.GetUuid(),
		Username: u.GetUsername(),
	}, nil
}
