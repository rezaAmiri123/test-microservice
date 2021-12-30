package agent

import (
	"github.com/go-playground/validator/v10"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
)

func (a *Agent) setupValidator() error {
	domain.SetValidator(validator.New())
	return nil
}
