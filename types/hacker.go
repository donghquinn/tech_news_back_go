package types

import "time"

type HackerNewsRequest struct {
	Today string `json:"today" xml:"today" binding:"required"`
}

type HackerNewsResponse struct {
	Uuid string `json:"uuid" xml:"uuid" binding:"required"`
	Post string `json:"post" xml:"post" binding:"required"`
	Link string `json:"link" xml:"link" binding:"required"`
	Founded time.Time `json:"founded" xml:"founded" binding:"required"`
}