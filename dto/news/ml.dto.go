package dto

import (
	"encoding/json"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/types"
	responseDto "github.com/dongquinn/tech_news_back_go/types/news"
)


func SetMlResponse(response http.ResponseWriter, result bool, code string, news []types.MachineLEarningNewsResponse) {
	bodyData ,_ := json.Marshal(responseDto.MlNewsResponseType{Result: result, Code: code, News: news})

	response.WriteHeader(200)
	response.Write(bodyData)
}

func SetMlErrorResponse(response http.ResponseWriter, result bool, code string, message string) {
	bodyData ,_ := json.Marshal(responseDto.MlNewsErrorResponseType{Result: result, Code: code, Message: message})

	response.WriteHeader(200)
	response.Write(bodyData)
}