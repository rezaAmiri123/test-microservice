package grpc_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	libraryservice "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"github.com/stretchr/testify/require"
)

func TestArticleGRPCServer_GetArticleBySlug(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testServer := NewTestGrpcServer(t, ctrl)
	req := &libraryservice.GetArticleBySlugRequest{
		Slug: "test_1",
	}
	a := &article.Article{
		Slug: req.GetSlug(),
	}
	testServer.repoMock.EXPECT().GetBySlug(gomock.Any(), gomock.Any()).Return(a, nil)
	ctx := context.Background()
	res, err := testServer.grpcServer.GetArticleBySlug(ctx, req)
	require.NoError(t, err)
	require.Equal(t, res.GetArticle().Slug, req.GetSlug())

}
