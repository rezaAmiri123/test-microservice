package dto

import (
	messageapi "github.com/rezaAmiri123/test-microservice/message_service/proto/grpc"
)

type GetEmailsResponse struct {
	List
	Emails []*EmailResponse `json:"emails"`
}

func EmailResponseFromGrpc(a *messageapi.Email) *EmailResponse {
	res := &EmailResponse{}
	res.Subject = a.GetSubject()
	res.To = a.GetTo()
	res.From = a.GetFrom()
	res.Body = a.GetBody()
	return res
}
