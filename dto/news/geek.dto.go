package dto

import (
	"encoding/json"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/types"
	responseDto "github.com/dongquinn/tech_news_back_go/types/news"
)


func SetGeekResponse(response http.ResponseWriter, result bool, code string, news []types.GeekNewsResponse) {
	bodyData ,_ := json.Marshal(responseDto.GeekNewsResponseType{Result: result, Code: code, News: news})

	response.WriteHeader(200)
	response.Write(bodyData)
}

func SetGeekErrorResponse(response http.ResponseWriter, result bool, code string, message string) {
	bodyData ,_ := json.Marshal(responseDto.GeekNewsErrorResponseType{Result: result, Code: code, Message: message})

	response.WriteHeader(200)
	response.Write(bodyData)
}