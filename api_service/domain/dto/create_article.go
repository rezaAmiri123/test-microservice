package dto

type CreateArticleRequest struct {
	ArticleID   string `json:"article-id" validate:"required"`
	UserID      string `json:"user-id" validate:"required"`
	Title       string `json:"title" validate:"required,min=3,max=250"`
	Description string `json:"description"`
	Body        string `json:"body" validate:"required,min=3,max=250"`
}

type CreateArticleResponse struct {
	ArticleID string `json:"article-id" validate:"required"`
}
