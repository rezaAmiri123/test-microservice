package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	messageservice "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *MessageGRPCServer) GetEmails(ctx context.Context, req *messageservice.GetEmailsRequest) (*messageservice.GetEmailsResponse, error) {
	//s.cfg.Metric.GetArticlesGrpcRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(ctx, "MessageGRPCServer.GetEmails")
	defer span.Finish()

	pq := pagnation.NewPaginationQuery(int(req.GetSize()), int(req.GetPage()))
	a, err := s.cfg.App.Queries.GetEmails.Handle(ctx, pq)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "email not found")
	}

	res := &messageservice.GetEmailsResponse{}
	res.Size = a.Size
	res.Page = a.Page
	res.TotalPages = a.TotalPages
	res.TotalCount = a.TotalCount
	res.HasMore = a.HasMore

	emailList := make([]*messageservice.Email, 0, len(a.Emails))
	for _, art := range a.Emails {
		emailList = append(emailList, email.EmailResponseToGrpc(art))
	}
	res.Emails = emailList

	return res, nil
}
