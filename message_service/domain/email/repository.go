//go:generate mockgen -source repository.go -destination mock/repository.go -package mock
package email

import "context"

type Repository interface {
	Create(ctx context.Context, email *Email) error
	GetByUUID(ctx context.Context, uuid string) (*Email, error)
}
