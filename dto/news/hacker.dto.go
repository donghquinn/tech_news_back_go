package dto

import (
	"encoding/json"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/types"
	responseDto "github.com/dongquinn/tech_news_back_go/types/news"
)


func SetHackerResponse(response http.ResponseWriter, result bool, code string, news []types.HackerNewsResponse) {
	bodyData ,_ := json.Marshal(responseDto.HackerNewsResponseType{Result: result, Code: code, News: news})

	response.WriteHeader(200)
	response.Write(bodyData)
}

func SetHackerErrorResponse(response http.ResponseWriter, result bool, code string, message string) {
	bodyData ,_ := json.Marshal(responseDto.HackerNewsErrorResponseType{Result: result, Code: code, Message: message})

	response.WriteHeader(200)
	response.Write(bodyData)
}