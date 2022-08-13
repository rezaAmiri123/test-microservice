package email

import (
	"context"
	"time"

	"github.com/google/uuid"
	messageapi "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/utils"
)

type List struct {
	TotalCount int64 `json:"total_count"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasMore    bool  `json:"has_more"`
}

type Email struct {
	UUID      string    `json:"uuid"`
	From      string    `json:"from" validate:"required,email"`
	To        []string  `json:"to" validate:"required"`
	Subject   string    `json:"subject" validate:"required,min=3,max=250"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EmailList struct {
	List
	Emails []*Email
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

func EmailResponseToGrpc(e *Email) *messageapi.Email {
	res := &messageapi.Email{}
	res.Subject = e.Subject
	res.To = e.To
	res.From = e.From
	res.Body = e.Body
	return res
}
