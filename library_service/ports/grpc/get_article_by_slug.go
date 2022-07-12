package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	libraryservice "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *articleGRPCServer) GetArticle(ctx context.Context, req *libraryservice.GetArticleRequest) (*libraryservice.GetArticleResponse, error) {
	s.cfg.Metric.GetArticleBySlugGrpcRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(ctx, "articleGRPCServer.GetArticle")
	defer span.Finish()

	a, err := s.cfg.App.Queries.GetArticleBySlug.Handle(ctx, req.GetSlug())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "article not found")
	}
	article := &libraryservice.Article{
		Title:       a.Title,
		Body:        a.Body,
		Description: a.Body,
		Slug:        a.Slug,
	}
	return &libraryservice.GetArticleResponse{Article: article}, nil
}
