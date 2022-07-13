package article

import (
	"context"
	"github.com/google/uuid"
	libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/utils"
)

type List struct {
	TotalCount int64 `json:"total_count"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasMore    bool  `json:"has_more"`
}

// Article model
type Article struct {
	UUID        string `json:"uuid"`
	UserUUID    string `json:"user_uuid"`
	Title       string `json:"title" validate:"required,min=3,max=250"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body" validate:"required,min=3,max=250"`
}

func ArticleResponseToGrpc(a *Article) *libraryapi.Article {
	res := &libraryapi.Article{}
	res.Title = a.Title
	res.Body = a.Body
	res.Description = a.Description
	return res
}

type ArticleList struct {
	List
	Articles []*Article
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
