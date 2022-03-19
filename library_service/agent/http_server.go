package agent

import (
	"fmt"
	"github.com/rezaAmiri123/test-microservice/library_service/ports/http"
)

func (a *Agent) setupHttpServer() error {
	httpAddress := fmt.Sprintf("%s:%d", a.Config.HttpServerAddr, a.Config.HttpServerPort)
	echoServer, err := http.NewHttpServer(httpAddress, a.Application, a.metric, a.AuthClient)
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
