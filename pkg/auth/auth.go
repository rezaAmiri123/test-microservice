package auth

import (
	"context"
)

type AuthClient interface {
	Login(ctx context.Context, username, password string) (string, error)
	VerifyToken(ctx context.Context, token string) (*User, error)
}

type User struct {
	Username string
	UUID     string
}

//
//func NewUserAuthClient(authClient UserApi.UsersServiceClient) (AuthClient, error) {
//	return &UserAuthClient{AuthClient: authClient}, nil
//}
//
//type UserAuthClient struct {
//	AuthClient UserApi.UsersServiceClient
//}
//
//func (a *UserAuthClient) Login(ctx context.Context, username, password string) (string, error) {
//	return "", nil
//}
//
//func (a *UserAuthClient) VerifyToken(ctx context.Context, token string) (*User, error) {
//	u, err := a.AuthClient.VerifyToken(ctx, &UserApi.VerifyTokenRequest{Token: token})
//	if err != nil {
//		return nil, err
//	}
//	return &User{
//		UUID:     u.GetUuid(),
//		Username: u.GetUsername(),
//	}, nil
//}
