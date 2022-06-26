package agent

import (
	"github.com/rezaAmiri123/test-microservice/pkg/logger/applogger"
)

func (a *Agent) setupLogger() error {
	appLogger := applogger.NewAppLogger(a.LoggerConfig)
	appLogger.InitLogger()
	appLogger.WithName("APIService")
	a.logger = appLogger
	return nil
}
