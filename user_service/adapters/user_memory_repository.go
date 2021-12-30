package adapters

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
	"sync"
)

type MemoryUserRepository struct {
	users map[string]domain.User
	mu    sync.RWMutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: map[string]domain.User{},
		mu:    sync.RWMutex{},
	}
}

func (m *MemoryUserRepository) Create(ctx context.Context, user *domain.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.users[user.Username]=*user
	return nil
}

func (m *MemoryUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	u, ok := m.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	return &u, nil
}
