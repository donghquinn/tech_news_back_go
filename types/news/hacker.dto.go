package types

import "github.com/dongquinn/tech_news_back_go/types"


type HackerNewsResponseType struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	News []types.HackerNewsResponse
}

type HackerNewsErrorResponseType struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Message string `json:"message"`
}