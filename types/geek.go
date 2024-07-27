package types

import "time"

type GeekNewsRequest struct {
	Today string `json:"today"`
}

type GeekNewsResponse struct {
	OriginalLink string `json:"originalLink"`
	DescLink string `json:"descLink"`
	Uuid string `json:"uuid"`
	Post  string `json:"post"`
	Founded time.Time `json:"founded"`
}