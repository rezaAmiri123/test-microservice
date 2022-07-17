package v1_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/converter"
	"github.com/stretchr/testify/require"
)

func TestHttpServer_CreateArticle(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testServer := NewTestHttpServer(t, ctrl)
	articleRequest := &dto.CreateArticleRequest{
		Title:       "title_1",
		Body:        "body_1",
		Description: "description_1",
	}
	buf, err := converter.AnyToBytesBuffer(articleRequest)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/articles/create", strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "bearer fffffffffffffffffffffff")
	res := httptest.NewRecorder()

	u := &auth.User{
		Username: "username_1",
		UUID:     "uuid_1",
	}

	testServer.authClientMock.EXPECT().VerifyToken(gomock.Any(), gomock.Any()).Return(u, nil)
	testServer.producerMock.EXPECT().PublishMessage(gomock.Any(), gomock.Any()).Return(nil)

	testServer.echoServer.ServeHTTP(res, req)
	require.NotNil(t, res.Body)
}
