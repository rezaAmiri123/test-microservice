package dto

type GetArticleBySlugRequest struct {
	Slug string `json:"slug" validate:"required,min=3,max=250"`
}

type GetArticleBySlugResponse struct {
	ArticleID   string `json:"article-id" validate:"required"`
	Title       string `json:"title" validate:"required,min=3,max=250"`
	Description string `json:"description"`
	Body        string `json:"body" validate:"required,min=3,max=250"`
}
