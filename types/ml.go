package types

import "time"

type MachineLearningNewsRequest struct {
	Today string `json:"today"`
}

type MachineLEarningNewsResponse struct {
	Uuid string `json:"uuid"`
	Category string `json:"category"`
	Title string `json:"title"`
	Link string `json:"link"`
	Founded time.Time `json:"founded"`
}