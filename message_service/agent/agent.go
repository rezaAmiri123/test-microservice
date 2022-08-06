package agent

import (
	"crypto/tls"
	"io"
	"net/http"
	"sync"

	"github.com/rezaAmiri123/test-microservice/message_service/app"
	"github.com/rezaAmiri123/test-microservice/message_service/metrics"
	kafkatopics "github.com/rezaAmiri123/test-microservice/message_service/ports/kafka"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/test-microservice/pkg/mongodb"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"google.golang.org/grpc"
)

type Config struct {
	ServerTLSConfig *tls.Config

	GRPCServerAddr string
	GRPCServerPort int
	//DBConfig     adapters.GORMConfig
	DBConfig       postgres.Config
	MongoConfig    mongodb.Config
	LoggerConfig   applogger.Config
	TracerConfig   tracing.Config
	MetricConfig   metrics.Config
	KafkaConfig    kafkatopics.Config
	RabbitmqConfig rabbitmq.Config
}

type Agent struct {
	Config

	logger      logger.Logger
	metric      *metrics.MessageServiceMetric
	httpServer  *http.Server
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
		a.setupKafka,
		//a.setupAuthClient,
		//a.setupHttpServer,
		a.setupGrpcServer,
		a.setupRabbitMQ,
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
		//func() error {
		//	return a.httpServer.Shutdown(context.Background())
		//},
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
