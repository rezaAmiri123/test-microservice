package email

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/pkg/utils"
)

type Email struct {
	UUID    string `json:"uuid"`
	From    string `json:"from" validate:"required, email"`
	To      string `json:"to" validate:"required, email"`
	Subject string `json:"subject" validate:"required,min=3,max=250"`
	Body    string `json:"body" validate:"required,min=3,max=250"`
}

// Validate validates fields of user model
func (e Email) Validate(ctx context.Context) error {
	if err := utils.ValidateStruct(ctx, e); err != nil {
		return err
	}
	return nil
}
