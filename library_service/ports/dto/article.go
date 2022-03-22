package dto

import (
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
)

type CreateArticleRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=250"`
	Description string `json:"description"`
	Body        string `json:"body" validate:"required,min=3,max=250"`
}

type CreateArticleResponse struct {
	Article article.Article
}

func (a CreateArticleRequest) MapToArticle() *article.Article {
	return &article.Article{
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
	}
}
