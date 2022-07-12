package agent

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (a *Agent) getLibraryClient() (grpc.ClientConnInterface, error) {
	// addr := fmt.Sprintf("%s:%d", config.GRPCUserAddr, config.GRPCUserPort)
	addr := fmt.Sprintf("%s:%d", a.GRPCLibraryClientAddr, a.GRPCLibraryClientPort)
	var opts []grpc.DialOption
	if a.GRPCLibraryClientTLSConfig != nil {
		clientCreds := credentials.NewTLS(a.GRPCLibraryClientTLSConfig)
		opts = append(opts, grpc.WithTransportCredentials(clientCreds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	//conn, err := grpc.Dial(addr, opts...)
	//if err != nil {
	//	return nil, err
	//}
	//authClient := libraryapi.NewArticleServiceClient(conn)
	//userAuthClient, _ := auth.NewUserAuthClient(authClient)
	//a.AuthClient = userAuthClient
	return grpc.Dial(addr, opts...)
}
