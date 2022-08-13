//go:generate mockgen -source repository.go -destination mock/repository.go -package mock
package email

import (
	"context"

	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
)

type Repository interface {
	Create(ctx context.Context, email *Email) error
	GetByUUID(ctx context.Context, uuid string) (*Email, error)
	List(ctx context.Context, query *pagnation.Pagination) (*EmailList, error)
}
