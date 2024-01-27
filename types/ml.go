package types

type MachineLearningNewsRequest struct {
	Today string `json:"today" xml:"today" binding:"required"`
}