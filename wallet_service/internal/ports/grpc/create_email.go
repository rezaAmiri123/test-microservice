package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	messageservice "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *MessageGRPCServer) CreateEmail(ctx context.Context, req *messageservice.CreateEmailRequest) (*messageservice.CreateEmailResponse, error) {
	//s.cfg.Metric.GetArticleBySlugGrpcRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(ctx, "MessageGRPCServer.CreateEmail")
	defer span.Finish()
	e := req.GetEmail()
	reqEmail := &email.Email{
		From:    e.GetFrom(),
		To:      e.GetTo(),
		Subject: e.GetSubject(),
		Body:    e.GetBody(),
	}
	resEmail, err := s.cfg.App.Commands.CreateEmailWithQueue.Handle(ctx, reqEmail)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &messageservice.CreateEmailResponse{UUID: resEmail.UUID}, nil
}
