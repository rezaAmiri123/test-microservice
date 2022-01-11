package domain_test

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		Name      string
		GetUserFn func() *domain.User
		HasError  bool
	}{
		{
			Name: "valid_user",
			GetUserFn: func() *domain.User {
				return newUser(t)
			},
			HasError: false,
		},
		{
			Name: "invalid_empty_email",
			GetUserFn: func() *domain.User {
				u := newUser(t)
				u.Email = ""
				return u
			},
			HasError: true,
		},
		{
			Name: "invalid_email",
			GetUserFn: func() *domain.User {
				u := newUser(t)
				u.Email = "123456789"
				return u
			},
			HasError: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			u := tc.GetUserFn()
			err := u.Validate(context.Background())
			if tc.HasError {
				require.True(t, err != nil)
			} else {
				require.True(t, err == nil)
			}
		})
	}
}

func newUser(t testing.TB) *domain.User {
	t.Helper()
	user := &domain.User{
		Username: "username",
		Password: "password",
		Email:    "email@example.com",
	}
	return user
}
