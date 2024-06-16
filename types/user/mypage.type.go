package types

import "github.com/dongquinn/tech_news_back_go/types"

type MyPageRequest struct {

}

type MyPageResponse struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	HackerNews []types.HackerNewsResponse	`json:"hackerNews"`
	GeekNews []types.GeekNewsResponse	`json:"geekNews"`
	MlNews []types.MachineLEarningNewsResponse	`json:"mlNews"`
}


type MyPageErrorResponse struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Message string `json:"message"`
}