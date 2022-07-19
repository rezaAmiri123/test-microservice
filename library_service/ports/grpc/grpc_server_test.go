package grpc_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/app/queries"
	repomock "github.com/rezaAmiri123/test-microservice/library_service/domain/article/mock"
	"github.com/rezaAmiri123/test-microservice/library_service/metrics"
	grpcervice "github.com/rezaAmiri123/test-microservice/library_service/ports/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/stretchr/testify/require"
)

type TestGrpcServer struct {
	grpcServer *grpcervice.ArticleGRPCServer
	logger     logger.Logger
	repoMock   *repomock.MockRepository
	Metric     *metrics.LibraryServiceMetric
}

func NewTestGrpcServer(t *testing.T, ctrl *gomock.Controller) *TestGrpcServer {
	t.Helper()

	appLogger := getLogger()

	repoMock := repomock.NewMockRepository(ctrl)

	application := &app.Application{
		Queries: app.Queries{
			GetArticleBySlug: queries.NewGetArticleBySlugHandler(repoMock),
		},
	}
	metric := metrics.NewLibraryServiceMetric(&metrics.Config{
		ServiceName: fmt.Sprintf("rand%d", rand.Int()),
	})

	serverConfig := grpcervice.Config{App: application, Metric: metric}
	server, err := grpcervice.NewArticleGRPCServer(&serverConfig)
	require.NoError(t, err)

	testServer := &TestGrpcServer{
		logger:     appLogger,
		Metric:     metric,
		repoMock:   repoMock,
		grpcServer: server,
	}
	return testServer
}

func getLogger() logger.Logger {
	appLogger := applogger.NewAppLogger(applogger.Config{})
	appLogger.InitLogger()
	appLogger.WithName("LibraryServiceTest")
	return appLogger
}
