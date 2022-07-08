package agent

import (
	"fmt"

	auth "github.com/rezaAmiri123/test-microservice/pkg/auth/auth_client"
	UserApi "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (a *Agent) setupAuthClient() error {
	// addr := fmt.Sprintf("%s:%d", config.GRPCUserAddr, config.GRPCUserPort)
	addr := fmt.Sprintf("%s:%d", a.GRPCAuthClientAddr, a.GRPCAuthClientPort)
	var opts []grpc.DialOption
	if a.GRPCAuthClientTLSConfig != nil {
		clientCreds := credentials.NewTLS(a.GRPCAuthClientTLSConfig)
		opts = append(opts, grpc.WithTransportCredentials(clientCreds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return err
	}
	authClient := UserApi.NewUsersServiceClient(conn)
	userAuthClient, _ := auth.NewUserAuthClient(authClient)
	a.AuthClient = userAuthClient
	return nil
}
