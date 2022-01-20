package agent

import (
	"fmt"

	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	UserApi "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
	"google.golang.org/grpc"
)

func (a *Agent) setupAuthClient() error {
	// addr := fmt.Sprintf("%s:%d", config.GRPCUserAddr, config.GRPCUserPort)
	addr := fmt.Sprintf("%s:%d", a.GRPCAuthClientAddr, a.GRPCAuthClientPort)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return err
	}
	authClient := UserApi.NewUsersServiceClient(conn)
	userAuthClient, _ := auth.NewUserAuthClient(authClient)
	a.AuthClient = userAuthClient
	return nil
}
