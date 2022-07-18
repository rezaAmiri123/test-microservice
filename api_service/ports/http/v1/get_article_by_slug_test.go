package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"github.com/stretchr/testify/require"
)

func TestHttpServer_GetArticleBySlug(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testServer := NewTestHttpServer(t, ctrl)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/articles/article/8a3cc26-fbe1-4713-98be-a2927201356e", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//req.Header.Set("Authorization", "bearer fffffffffffffffffffffff")
	res := httptest.NewRecorder()

	articleBySlugResponse := &libraryapi.GetArticleBySlugResponse{
		Article: &libraryapi.Article{
			Body:        "body_",
			Title:       "title_",
			Description: "description_",
			Slug:        "8a3cc26-fbe1-4713-98be-a2927201356e",
		},
	}
	testServer.articleClientMock.SetArticleBySlugResponse(articleBySlugResponse)

	testServer.echoServer.ServeHTTP(res, req)
	require.NotNil(t, res.Body)
}
