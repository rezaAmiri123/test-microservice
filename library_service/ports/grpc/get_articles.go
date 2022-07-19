package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	libraryservice "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ArticleGRPCServer) GetArticles(ctx context.Context, req *libraryservice.GetArticlesRequest) (*libraryservice.GetArticlesResponse, error) {
	s.cfg.Metric.GetArticlesGrpcRequests.Inc()

	span, ctx := opentracing.StartSpanFromContext(ctx, "articleGRPCServer.GetArticles")
	defer span.Finish()

	pq := pagnation.NewPaginationQuery(int(req.GetSize()), int(req.GetPage()))
	a, err := s.cfg.App.Queries.GetArticles.Handle(ctx, pq)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "article not found")
	}

	res := &libraryservice.GetArticlesResponse{}
	res.Size = a.Size
	res.Page = a.Page
	res.TotalPages = a.TotalPages
	res.TotalCount = a.TotalCount
	res.HasMore = a.HasMore

	articleList := make([]*libraryservice.Article, 0, len(a.Articles))
	for _, art := range a.Articles {
		articleList = append(articleList, article.ArticleResponseToGrpc(art))
	}
	res.Articles = articleList

	return res, nil
}
