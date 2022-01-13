package grpc

import (
	"context"
	userservice "github.com/rezaAmiri123/test-microservice/user_service/proto/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcServer) Login(ctx context.Context, req *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	u, err := s.cfg.App.Queries.GetProfile.Handle(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	token, err := u.GenerateJWTToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can not generate token")
	}
	return &userservice.LoginResponse{Token: token}, nil
}

func (s *grpcServer) VerifyToken(ctx context.Context, req *userservice.VerifyTokenRequest) (*userservice.User, error) {
	u, err := s.cfg.App.Queries.GetUserToken.Handler(ctx, req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &userservice.User{
		Username: u.Username,
		Uuid:     u.UUID,
	}, nil
}
