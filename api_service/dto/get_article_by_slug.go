package dto

type GetArticleBySlugRequest struct {
	Slug string `json:"slug" validate:"required,min=3,max=250"`
}

type GetArticleBySlugResponse struct {
	ArticleResponse
}
