package agent

import (
	"crypto/tls"
	"io"
	"net/http"
	"sync"

	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/app"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/domain/wallet"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/metrics"
	"google.golang.org/grpc"
)

type Config struct {
	ServerTLSConfig *tls.Config

	GRPCServerAddr string `mapstructure:"GRPC_SERVER_ADDR"`
	GRPCServerPort int    `mapstructure:"GRPC_SERVER_PORT"`
	//DBConfig     adapters.GORMConfig
	DBConfig     postgres.Config
	LoggerConfig applogger.Config
	TracerConfig tracing.Config
	MetricConfig metrics.Config
}

type Agent struct {
	Config

	logger      logger.Logger
	metric      *metrics.WalletServiceMetric
	httpServer  *http.Server
	grpcServer  *grpc.Server
	repository  wallet.Repository
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
		//a.setupAuthClient,
		//a.setupHttpServer,
		a.setupGrpcServer,
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
