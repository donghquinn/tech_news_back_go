package types

import "github.com/dongquinn/tech_news_back_go/types"


type MlNewsResponseType struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	News []types.MachineLEarningNewsResponse
}

type MlNewsErrorResponseType struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Message string `json:"message"`
}