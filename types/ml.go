package types

import "time"

type MachineLearningNewsRequest struct {
	Today string `json:"today" xml:"today" binding:"required"`
}

type MachineLEarningNewsResponse struct {
	Uuid string `json:"uuid" xml:"uuid" binding:"required"`
	Category string `json:"category" xml:"category" binding:"required"`
	Title string `json:"title" xml:"title" binding:"required"`
	Link string `json:"link" xml:"link" binding:"required"`
	Founded time.Time `json:"founded" xml:"founded" binding:"required"`
}