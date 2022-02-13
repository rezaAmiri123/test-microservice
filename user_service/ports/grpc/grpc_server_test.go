package grpc_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/rezaAmiri123/test-microservice/pkg/auth/tls"
	"github.com/rezaAmiri123/test-microservice/user_service/app"
	"github.com/rezaAmiri123/test-microservice/user_service/app/command"
	"github.com/rezaAmiri123/test-microservice/user_service/app/query"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"github.com/rezaAmiri123/test-microservice/user_service/domain/mock"
	server "github.com/rezaAmiri123/test-microservice/user_service/ports/grpc"
	serverproto "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
	"github.com/stretchr/testify/require"
	"github.com/travisjeffery/go-dynaport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"testing"
)

var (
	CAFile         = tls.ConfigFile("ca.pem")
	ServerCertFile = tls.ConfigFile("server.pem")
	ServerKeyFile  = tls.ConfigFile("server-key.pem")
	ClientCertFile = tls.ConfigFile("client.pem")
	ClientKeyFile  = tls.ConfigFile("client-key.pem")
)

func TestGRPCServer(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		client serverproto.UsersServiceClient,
		config *server.Config,
		repo *mock.MockRepository,
	){
		"login test":   testGrpcLoginUser,
		"verify token": testGrpcServer_VerifyTokenUser,
	} {
		t.Run(scenario, func(t *testing.T) {
			client, config, repo, teardown := setupGRPCServerTest(t, nil)
			defer teardown()
			fn(t, client, config, repo)
		})
	}
}

func setupGRPCServerTest(t *testing.T, fn func(config *server.Config)) (
	client serverproto.UsersServiceClient,
	cfg *server.Config,
	repo *mock.MockRepository,
	teardown func(),
) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo = mock.NewMockRepository(ctrl)
	application := &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(repo, nil, nil),
		},
		Queries: app.Queries{
			GetProfile:   query.NewGetProfileHandler(repo),
			GetUserToken: query.NewGetUserTokenHandler(repo),
		},
	}

	httpPorts := dynaport.Get(1)
	bindAddr := fmt.Sprintf("%s:%d", "127.0.0.1", httpPorts[0])
	serverConfig := server.Config{application}
	var opts []grpc.ServerOption
	serverTLSConfig, err := tls.SetupTLSConfig(tls.TLSConfig{
		CAFile:   CAFile,
		CertFile: ServerCertFile,
		KeyFile:  ServerKeyFile,
		Server:   true,
	})
	require.NoError(t, err)
	serverCreds := credentials.NewTLS(serverTLSConfig)
	opts = append(opts, grpc.Creds(serverCreds))
	grpcServer, err := server.NewGRPCServer(&serverConfig, opts...)
	ln, err := net.Listen("tcp", bindAddr)
	require.NoError(t, err)
	go func() {
		grpcServer.Serve(ln)
	}()
	clientTLSCofig, err := tls.SetupTLSConfig(tls.TLSConfig{
		CAFile:   CAFile,
		CertFile: ClientCertFile,
		KeyFile:  ClientKeyFile,
	})
	require.NoError(t, err)
	clientCreds := credentials.NewTLS(clientTLSCofig)

	clientOptions := []grpc.DialOption{
		//grpc.WithInsecure(),
		grpc.WithTransportCredentials(clientCreds),
	}
	cc, err := grpc.Dial(bindAddr, clientOptions...)
	require.NoError(t, err)
	client = serverproto.NewUsersServiceClient(cc)
	return client, cfg, repo, func() {
		grpcServer.Stop()
		cc.Close()
		ln.Close()
	}
}

func testGrpcLoginUser(t *testing.T, client serverproto.UsersServiceClient, config *server.Config, repo *mock.MockRepository) {
	ctx := context.Background()
	user := &domain.User{
		Username: "username",
		Password: "password",
		Email:    "email@example.com",
	}
	//err := application.Commands.CreateUser.Handle(ctx, want)
	//require.NoError(t,err)
	repo.EXPECT().GetByUsername(gomock.Any(), user.Username).Return(user, nil)
	token, err := client.Login(ctx, &serverproto.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})
	require.NoError(t, err)
	require.NotEqual(t, token, "")
}

func testGrpcServer_VerifyTokenUser(t *testing.T, client serverproto.UsersServiceClient, config *server.Config, repo *mock.MockRepository) {
	ctx := context.Background()
	user := &domain.User{
		Username: "username",
		Password: "password",
		Email:    "email@example.com",
	}
	//err := application.Commands.CreateUser.Handle(ctx, want)
	//require.NoError(t,err)
	repo.EXPECT().GetByUsername(gomock.Any(), user.Username).Return(user, nil)
	token, err := client.Login(ctx, &serverproto.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})
	require.NoError(t, err)
	repo.EXPECT().GetByUsername(gomock.Any(), user.Username).Return(user, nil)
	got, err := client.VerifyToken(ctx, &serverproto.VerifyTokenRequest{
		Token: token.GetToken(),
	})
	require.NoError(t, err)
	require.Equal(t, got.Username, user.Username)
}
