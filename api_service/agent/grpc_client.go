package agent

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc/credentials"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/rezaAmiri123/test-microservice/pkg/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	backoffLinear  = 100 * time.Millisecond
	backoffRetries = 3
)

func (a *Agent) getGrpcClient(addr string, clientTLSConfig *tls.Config) (grpc.ClientConnInterface, error) {
	// addr := fmt.Sprintf("%s:%d", config.GRPCUserAddr, config.GRPCUserPort)
	im := interceptors.NewInterceptorManager(a.logger)
	//addr := fmt.Sprintf("%s:%d", a.GRPCLibraryClientAddr, a.GRPCLibraryClientPort)
	ctx := context.Background()

	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backoffLinear)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
		grpc_retry.WithMax(backoffRetries),
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithChainUnaryInterceptor(
		grpc_retry.UnaryClientInterceptor(retryOpts...),
		im.ClientRequestLoggerInterceptor(),
	))

	//if a.GRPCLibraryClientTLSConfig != nil {
	//	clientCreds := credentials.NewTLS(a.GRPCLibraryClientTLSConfig)
	//	opts = append(opts, grpc.WithTransportCredentials(clientCreds))
	//} else {
	//	opts = append(opts, grpc.WithInsecure())
	//}
	if clientTLSConfig != nil {
		clientCreds := credentials.NewTLS(clientTLSConfig)
		opts = append(opts, grpc.WithTransportCredentials(clientCreds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	return grpc.DialContext(ctx, addr, opts...)
}
