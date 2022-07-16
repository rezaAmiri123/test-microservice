package v1_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/api_service/app"
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	"github.com/rezaAmiri123/test-microservice/api_service/metrics"
	v1 "github.com/rezaAmiri123/test-microservice/api_service/ports/http/v1"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	authMock "github.com/rezaAmiri123/test-microservice/pkg/auth/mock"
	"github.com/rezaAmiri123/test-microservice/pkg/converter"
	kafkaMock "github.com/rezaAmiri123/test-microservice/pkg/kafka/mock"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/stretchr/testify/require"
)

func getLogger() logger.Logger {
	appLogger := applogger.NewAppLogger(applogger.Config{})
	appLogger.InitLogger()
	appLogger.WithName("APIServiceTest")
	return appLogger
}

func TestHttpServer_CreateArticle(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	appLogger := getLogger()
	producerMock := kafkaMock.NewMockProducer(ctrl)
	authClientMock := authMock.NewMockAuthClient(ctrl)
	application := &app.Application{Commands: app.Commands{
		CreateArticle: command.NewCreateArticleHandler(producerMock, appLogger),
	}}
	metric := metrics.NewApiServiceMetric(&metrics.Config{})

	echoServer, err := v1.NewHttpServer(false, application, metric, appLogger, authClientMock)
	require.NoError(t, err)

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

	authClientMock.EXPECT().VerifyToken(gomock.Any(), gomock.Any()).Return(u, nil)
	producerMock.EXPECT().PublishMessage(gomock.Any(), gomock.Any()).Return(nil)

	echoServer.ServeHTTP(res, req)
	require.NotNil(t, res.Body)
}
