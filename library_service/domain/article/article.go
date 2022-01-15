package article

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/test-microservice/pkg/utils"
)

// Article model
type Article struct {
	UUID        string `json:"uuid"`
	UserUUID    string `json:"user_uuid"`
	Title       string `json:"title" validate:"required,min=3,max=250"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body" validate:"required,min=3,max=250"`
}

// Validate validates fields of user model
func (a Article) Validate(ctx context.Context) error {
	if err := utils.ValidateStruct(ctx, a); err != nil {
		return err
	}
	return nil
}

func (a *Article) SetUUID(userUUID string) error {
	if a.UUID == "" {
		a.UUID = uuid.New().String()
	}
	a.UserUUID = userUUID
	return nil
}
