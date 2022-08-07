package dto

type GetEmailByUUIDRequest struct {
	UUID string `json:"uuid" validate:"required,min=3,max=250"`
}

type GetEmailByUUIDResponse struct {
	EmailResponse
}

type EmailResponse struct {
	Subject string   `json:"subject" validate:"required"`
	To      []string `json:"to" validate:"required,min=3,max=250"`
	From    string   `json:"from"`
	Body    string   `json:"body" validate:"required,min=3,max=250"`
}
