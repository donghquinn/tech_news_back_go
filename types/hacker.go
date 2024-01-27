package types

type HackerNewsRequest struct {
	Today string `json:"today" xml:"today" binding:"required"`
}