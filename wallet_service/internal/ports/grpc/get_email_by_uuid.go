package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	messageservice "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *MessageGRPCServer) GetEmailByUUID(ctx context.Context, req *messageservice.GetEmailByUUIDRequest) (*messageservice.GetEmailByUUIDResponse, error) {
	//s.cfg.Metric.GetArticleBySlugGrpcRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(ctx, "articleGRPCServer.GetEmailByUUID")
	defer span.Finish()

	a, err := s.cfg.App.Queries.GetEmailByUUID.Handle(ctx, req.GetUUID())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "email not found")
	}
	res := &messageservice.GetEmailByUUIDResponse{
		Subject: a.Subject,
		To:      a.To,
		From:    a.From,
		Body:    a.Body,
	}
	return res, nil
}
