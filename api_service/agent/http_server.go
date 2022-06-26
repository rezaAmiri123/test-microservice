package agent

import (
	"fmt"
	"github.com/rezaAmiri123/test-microservice/api_service/ports/http/v1"
)

func (a *Agent) setupHttpServer() error {
	httpAddress := fmt.Sprintf("%s:%d", a.Config.HttpServerAddr, a.Config.HttpServerPort)
	echoServer, err := v1.NewHttpServer(a.Application, a.metric, a.logger)
	if err != nil {
		return err
	}
	a.httpServer = echoServer
	go func() {
		if err := a.httpServer.Start(httpAddress); err != nil {
			_ = a.Shutdown()
		}
	}()

	return nil
}
