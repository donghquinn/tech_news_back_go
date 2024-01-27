package types

import "time"

type MachineLearningNewsRequest struct {
	Today string `json:"today" xml:"today" binding:"required"`
}

type MachineLEarningNewsResponse struct {
	Uuid string `json:"uuid" xml:"uuid" binding:"required,min=1"`
	Category string `json:"category" xml:"category" binding:"required,min=1"`
	Title string `json:"title" xml:"title" binding:"required,min=1"`
	Link string `json:"link" xml:"link" binding:"required,min=1"`
	Founded time.Time `json:"founded" xml:"founded" binding:"required"`
}