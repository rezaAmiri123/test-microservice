package agent

import (
	"context"
	"crypto/tls"
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/pkg/db/postgres"
	"io"
	"sync"

	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/metrics"
	kafkatopics "github.com/rezaAmiri123/test-microservice/library_service/ports/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"google.golang.org/grpc"
)

type Config struct {
	HttpServerAddr          string
	HttpServerPort          int
	GRPCServerAddr          string
	GRPCServerPort          int
	GRPCAuthClientAddr      string
	GRPCAuthClientPort      int
	GRPCAuthClientTLSConfig *tls.Config

	//DBConfig     adapters.GORMConfig
	DBConfig     postgres.Config
	LoggerConfig applogger.Config
	TracerConfig tracing.Config
	MetricConfig metrics.Config
	KafkaConfig  kafkatopics.Config
}

type Agent struct {
	Config

	logger      logger.Logger
	metric      *metrics.LibraryServiceMetric
	httpServer  *echo.Echo
	grpcServer  *grpc.Server
	repository  domain.Repository
	Application *app.Application
	AuthClient  auth.AuthClient

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
	closers      []io.Closer
}

func NewAgent(config Config) (*Agent, error) {
	a := &Agent{
		Config:    config,
		shutdowns: make(chan struct{}),
	}
	setupsFn := []func() error{
		a.setupLogger,
		a.setupMetric,

		//a.setupRepository,
		a.setupTracing,
		a.setupApplication,
		a.setupAuthClient,
		a.setupHttpServer,
		a.setupKafka,
		//a.setupGrpcServer,
		//a.setupGRPCServer,
		//a.setupTracer,
	}
	for _, fn := range setupsFn {
		if err := fn(); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (a *Agent) Shutdown() error {
	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()

	if a.shutdown {
		return nil
	}
	a.shutdown = true
	close(a.shutdowns)
	shutdown := []func() error{
		func() error {
			return a.httpServer.Shutdown(context.Background())
		},
		func() error {
			a.grpcServer.GracefulStop()
			return nil
		},
		//func() error {
		//	return a.jaegerCloser.Close()
		//},
	}
	for _, fn := range shutdown {
		if err := fn(); err != nil {
			return err
		}
	}
	for _, closer := range a.closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}
