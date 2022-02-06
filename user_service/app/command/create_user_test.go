package command_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	kafkamock "github.com/rezaAmiri123/test-microservice/pkg/kafka/mock"
	"github.com/rezaAmiri123/test-microservice/user_service/app/command"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"github.com/rezaAmiri123/test-microservice/user_service/domain/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUserHandler_Handle(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockRepository(ctrl)
	producer := kafkamock.NewMockProducer(ctrl)
	createUsrHandler := command.NewCreateUserHandler(repo, producer, nil)
	user := &domain.User{
		Username: "username",
		Password: "password",
		Email:    "email@example.com",
	}
	ctx := context.Background()
	repo.EXPECT().Create(gomock.Any(), gomock.Eq(user)).Return(nil)
	producer.EXPECT().PublishMessage(gomock.Any(), gomock.Any()).Return(nil)
	err := createUsrHandler.Handle(ctx, user)
	require.NoError(t, err)
}
