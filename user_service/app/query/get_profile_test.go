package query_test

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/user_service/app/query"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"github.com/rezaAmiri123/test-microservice/user_service/domain/mock"
	"github.com/stretchr/testify/require"
)

func TestGetProfileHandler_Handle(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepository(ctrl)
	getProfileHandler := query.NewGetProfileHandler(repo)
	user := &domain.User{
		Username: "username",
		Password: "password",
		Email:    "email@example.com",
	}
	ctx := context.Background()
	repo.EXPECT().GetByUsername(gomock.Any(), user.Username).Return(user, nil)
	resp, err := getProfileHandler.Handle(ctx, user.Username)
	require.NoError(t, err)
	require.Equal(t, resp.Email, user.Email)

}
