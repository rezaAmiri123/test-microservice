//go:generate mockgen -source repository.go -destination mock/repository.go -package mock
package article

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
)

type Repository interface {
	Create(ctx context.Context, article *Article) error
	GetBySlug(ctx context.Context, slug string) (*Article, error)
	List(ctx context.Context, query *pagnation.Pagination) (*ArticleList, error)

	CreateComment(ctx context.Context, comment *Comment) error
}
