package agent

import (
	"fmt"
	"github.com/rezaAmiri123/test-microservice/user_service/ports"
)

func (a *Agent) setupHttpServer() error {
	httpAddress := fmt.Sprintf("%s:%d", a.Config.HttpServerAddr,a.Config.HttpServerPort)
	httpServer, err := ports.NewHttpServer(httpAddress, a.Application)
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
