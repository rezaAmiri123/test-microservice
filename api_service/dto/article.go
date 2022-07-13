package dto

import libraryapi "github.com/rezaAmiri123/test-microservice/library_service/proto/grpc"

type ArticleResponse struct {
	ArticleID   string `json:"article_id" validate:"required"`
	Title       string `json:"title" validate:"required,min=3,max=250"`
	Description string `json:"description"`
	Body        string `json:"body" validate:"required,min=3,max=250"`
}

func ArticleResponseFromGrpc(a *libraryapi.Article) *ArticleResponse {
	res := &ArticleResponse{}
	res.Title = a.GetTitle()
	res.Body = a.GetBody()
	res.Description = a.GetDescription()
	return res
}

type List struct {
	TotalCount int64 `json:"total_count"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasMore    bool  `json:"has_more"`
}
