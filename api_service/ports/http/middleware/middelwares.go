package middleware

import (
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	logger     logger.Logger
	authClient auth.AuthClient
	//origins    []string
}

// Middleware manager constructor
func NewMiddlewareManager(logger logger.Logger, authClient auth.AuthClient) *MiddlewareManager {
	return &MiddlewareManager{logger: logger, authClient: authClient}
}
