package dto

type GetArticlesResponse struct {
	List
	Articles []*ArticleResponse `json:"articles"`
}
