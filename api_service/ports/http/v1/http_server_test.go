package v1_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/api_service/app"
	"github.com/rezaAmiri123/test-microservice/api_service/app/command"
	"github.com/rezaAmiri123/test-microservice/api_service/metrics"
	v1 "github.com/rezaAmiri123/test-microservice/api_service/ports/http/v1"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	authMock "github.com/rezaAmiri123/test-microservice/pkg/auth/mock"
	kafkaMock "github.com/rezaAmiri123/test-microservice/pkg/kafka/mock"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type TestHttpServer struct {
	echoServer        *echo.Echo
	logger            logger.Logger
	producerMock      *kafkaMock.MockProducer
	authClientMock    *authMock.MockAuthClient
	articleClientMock libraryapi.ArticleServiceClient
}

func NewTestHttpServer(t *testing.T, ctrl *gomock.Controller) *TestHttpServer {
	t.Helper()

	appLogger := getLogger()
	producerMock := kafkaMock.NewMockProducer(ctrl)
	authClientMock := authMock.NewMockAuthClient(ctrl)
	application := &app.Application{Commands: app.Commands{
		CreateArticle: command.NewCreateArticleHandler(producerMock, appLogger),
	}}
	metric := metrics.NewApiServiceMetric(&metrics.Config{})

	echoServer, err := v1.NewHttpServer(false, application, metric, appLogger, authClientMock)
	require.NoError(t, err)

	testServer := &TestHttpServer{
		echoServer:        echoServer,
		logger:            appLogger,
		authClientMock:    authClientMock,
		producerMock:      producerMock,
		articleClientMock: &ArticleClientMock{},
	}
	return testServer
}

func getLogger() logger.Logger {
	appLogger := applogger.NewAppLogger(applogger.Config{})
	appLogger.InitLogger()
	appLogger.WithName("APIServiceTest")
	return appLogger
}

var _ libraryapi.ArticleServiceClient = (*ArticleClientMock)(nil)

type ArticleClientMock struct {
	ArticleBySlugResponse *libraryapi.GetArticleBySlugResponse
	ArticlesResponse      *libraryapi.GetArticlesResponse
	Err                   error
}

func (m *ArticleClientMock) SetArticleBySlugResponse(article *libraryapi.GetArticleBySlugResponse) {
	m.ArticleBySlugResponse = article
}
func (m *ArticleClientMock) SetArticlesResponse(article *libraryapi.GetArticlesResponse) {
	m.ArticlesResponse = article
}
func (m *ArticleClientMock) SetError(err error) {
	m.Err = err
}
func (m *ArticleClientMock) GetArticleBySlug(ctx context.Context, in *libraryapi.GetArticleBySlugRequest, opts ...grpc.CallOption) (*libraryapi.GetArticleBySlugResponse, error) {
	return m.ArticleBySlugResponse, m.Err
}

func (m *ArticleClientMock) GetArticles(ctx context.Context, in *libraryapi.GetArticlesRequest, opts ...grpc.CallOption) (*libraryapi.GetArticlesResponse, error) {
	return m.ArticlesResponse, m.Err
}
