package dto

import (
	"encoding/json"
	"net/http"

	types "github.com/dongquinn/tech_news_back_go/types/middleware"
)


func SetMiddlewareErrorResponse(response http.ResponseWriter,result bool, code string, message string) {
	bodyData, _ := json.Marshal(types.MiddlewareErrorResponse{Result: result, Code: code, Message: message})

	response.WriteHeader(200)
	response.Write(bodyData)
}