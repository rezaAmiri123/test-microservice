//go:generate mockgen -source repository.go -destination mock/repository.go -package mock
package domain

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	//Update(ctx context.Context, user *User)error
	//GetByEmail(ctx context.Context, email string)(*User,error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}
