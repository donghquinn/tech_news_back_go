package types

import "time"

type GeekNewsRequest struct {
	Today string `json:"today" xml:"today" binding:"required"`
}

type GeekNewsResponse struct {
	OriginalLink string `json:"originalLink" xml:"originalLink" binding:"required,min=1"`
	DescLink string `json:"descLink" xml:"descLink" binding:"required,min=1"`
	Uuid string `json:"uuid" xml:"uuid" binding:"required,min=1"`
	Post  string `json:"post" xml:"post" binding:"required,min=1"`
	Founded time.Time `json:"founded" xml:"founded" binding:"required"`
}