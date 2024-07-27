package types

import "github.com/dongquinn/tech_news_back_go/types"


type GeekNewsResponseType struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	News []types.GeekNewsResponse	`json:"news"`
}

type GeekNewsErrorResponseType struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Message string `json:"message"`
}