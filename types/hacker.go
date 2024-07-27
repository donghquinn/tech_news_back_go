package types

import "time"

type HackerNewsRequest struct {
	Today string `json:"today"`
}

type HackerNewsResponse struct {
	Uuid string `json:"uuid"`
	Rank string `json:"rank"`
	Post string `json:"post"`
	Link string `json:"link"`
	Founded time.Time `json:"founded"`
}