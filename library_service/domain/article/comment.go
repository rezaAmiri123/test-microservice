package article

import (
	"context"
	"time"

	"github.com/rezaAmiri123/test-microservice/pkg/utils"
)

// Comment model
type Comment struct {
	UUID        string    `json:"uuid"`
	UserUUID    string    `json:"user_uuid"`
	ArticleUUID string    `json:"article_uuid"`
	Article     *Article  `json:"article" validate:"excluded_with_all"`
	Message     string    `json:"message" validate:"required,gte=10"`
	Likes       int64     `json:"likes" validate:"omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate validates fields of user model
func (a Comment) Validate(ctx context.Context) error {
	if err := utils.ValidateStruct(ctx, a); err != nil {
		return err
	}
	return nil
}

//func (a *Comment) SetUUID(userUUID string) error {
//	if a.UUID == "" {
//		a.UUID = uuid.New().String()
//	}
//	a.UserUUID = userUUID
//	return nil
//}
