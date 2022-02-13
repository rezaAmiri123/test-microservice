package agent

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"net"

	grpcervice "github.com/rezaAmiri123/test-microservice/user_service/ports/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (a *Agent) setupGrpcServer() error {
	serverConfig := grpcervice.Config{App: a.Application}
	//serverConfig := ports.GRPCConfig{a.Application}
	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_auth.StreamServerInterceptor(auth.Authenticate),
			)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.Authenticate),
		)),
	)
	if a.Config.ServerTLSConfig != nil {
		creds := credentials.NewTLS(a.Config.ServerTLSConfig)
		opts = append(opts, grpc.Creds(creds))
	}
	var err error
	a.grpcServer, err = grpcervice.NewGRPCServer(&serverConfig, opts...)
	if err != nil {
		return err
	}
	grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCServerAddr, a.Config.GRPCServerPort)
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
