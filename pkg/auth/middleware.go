package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	UserApi "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
)

type UserHttpMiddleware struct {
	AuthClient UserApi.UsersServiceClient
}

func (m UserHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := m.tokenFromHeader(r)
		if bearerToken == "" {
			http.Error(w, "empty bearer token", http.StatusForbidden)
			return
		}
		tokenUser, err := m.AuthClient.VerifyToken(ctx, &UserApi.VerifyTokenRequest{Token: bearerToken})
		if err != nil {
			http.Error(w, "unable verify", http.StatusForbidden)
			return
		}
		ctx = context.WithValue(ctx, userContextKey, User{
			UUID:     tokenUser.GetUuid(),
			Username: tokenUser.GetUsername(),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func (m UserHttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	NoUserInContextError = errors.New("no user in context")
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
