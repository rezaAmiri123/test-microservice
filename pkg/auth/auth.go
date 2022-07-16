//go:generate mockgen -source auth.go -destination mock/auth.go -package mock
package auth

import (
	"context"
	"errors"
)

type ctxKey int

const (
	UserContextKey ctxKey = iota
)

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	NoUserInContextError = errors.New("no user found")
)

func UserFromCtx(ctx context.Context) *User {
	u, ok := ctx.Value(UserContextKey).(*User)
	if ok {
		return u
	}

	return &User{}
}

type User struct {
	Username string
	UUID     string
}

type AuthClient interface {
	Login(ctx context.Context, username, password string) (string, error)
	VerifyToken(ctx context.Context, token string) (*User, error)
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
