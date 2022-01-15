//go:generate mockgen -source repository.go -destination mock/repository.go -package mock
package article

import "context"

type Repository interface {
	Create(ctx context.Context, article *Article) error
	GetBySlug(ctx context.Context, slug string) (*Article, error)
}
