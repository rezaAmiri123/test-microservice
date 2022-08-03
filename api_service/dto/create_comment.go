package dto

type CreateCommentRequest struct {
	UUID        string `json:"uuid" validate:"required"`
	UserUUID    string `json:"user_uuid" validate:"required"`
	ArticleUUID string `json:"article_uuid" validate:"required"`
	Message     string `json:"message" validate:"required,min=3,max=250"`
}

type CreateCommentResponse struct {
	UUID string `json:"uuid" validate:"required"`
}
