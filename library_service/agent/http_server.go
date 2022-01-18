package agent

import (
	"fmt"
	"github.com/rezaAmiri123/test-microservice/library_service/ports/http"
)

func (a *Agent) setupHttpServer() error {
	httpAddress := fmt.Sprintf("%s:%d", a.Config.HttpServerAddr, a.Config.HttpServerPort)
	httpServer, err := http.NewHttpServer(httpAddress, a.Application, a.metric)
	if err != nil {
		return err
	}
	a.httpServer = httpServer
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			_ = a.Shutdown()
		}
	}()

	return nil
}
