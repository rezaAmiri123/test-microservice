package agent

import (
	"fmt"
	grpcervice "github.com/rezaAmiri123/test-microservice/user_service/ports/grpc"
	"google.golang.org/grpc"
	"net"
)

func (a *Agent) setupGrpcServer() error {
	serverConfig := grpcervice.Config{App: a.Application}
	//serverConfig := ports.GRPCConfig{a.Application}
	var opts []grpc.ServerOption
	var err error
	a.grpcServer, err = grpcervice.NewGRPCServer(&serverConfig, opts...)
	if err != nil {
		return err
	}
	grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCServerAddr,a.Config.GRPCServerPort)
	ln, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}
	go func() {
		if err := a.grpcServer.Serve(ln); err != nil {
			_ = a.Shutdown()
		}
	}()
	return err
}
