package agent

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
	"github.com/rezaAmiri123/test-microservice/user_service/adapters"
	"github.com/rezaAmiri123/test-microservice/user_service/app"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"io"
	"net/http"
	"sync"
)

type Config struct {
	HttpServerPort int
	HttpServerAddr string
	GRPCServerPort int
	GRPCServerAddr string

	DBConfig adapters.GORMConfig
	LoggerConfig applogger.Config
	TracerConfig tracing.Config
}

type Agent struct {
	Config

	logger      logger.Logger
	httpServer  *http.Server
	repository  domain.Repository
	Application *app.Application

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
	closers []io.Closer
}

func NewAgent(config Config) (*Agent, error) {
	a := &Agent{
		Config:    config,
		shutdowns: make(chan struct{}),
	}
	setupsFn := []func() error{
		a.setupLogger,
		//a.setupRepository,
		a.setupTracing,
		a.setupApplication,
		a.setupHttpServer,
		a.setupValidator,
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
		//func() error {
		//	a.grpcServer.GracefulStop()
		//	return nil
		//},
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
