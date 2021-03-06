package email

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/test-microservice/pkg/utils"
	"time"
)

type Email struct {
	UUID      string    `json:"uuid"`
	From      string    `json:"from" validate:"required,email"`
	To        []string  `json:"to" validate:"required"`
	Subject   string    `json:"subject" validate:"required,min=3,max=250"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validates fields of user model
func (e Email) Validate(ctx context.Context) error {
	if err := utils.ValidateStruct(ctx, e); err != nil {
		return err
	}
	return nil
}

func (e *Email) SetUUID() error {
	if e.UUID == "" {
		e.UUID = uuid.New().String()
	}
	return nil
}
