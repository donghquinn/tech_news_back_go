package types

import "time"

type HackerNewsRequest struct {
	Today string `json:"today" xml:"today" binding:"required"`
}

type HackerNewsResponse struct {
	Uuid string `json:"uuid" xml:"uuid" binding:"required,min=1"`
	Post string `json:"post" xml:"post" binding:"required,min=1"`
	Link string `json:"link" xml:"link" binding:"required,min=1"`
	Founded time.Time `json:"founded" xml:"founded" binding:"required"`
}