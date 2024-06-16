package dto

import (
	"encoding/json"
	"net/http"

	types "github.com/dongquinn/tech_news_back_go/types/user"
)

func SetLoginResponse(response http.ResponseWriter, result bool, code string, token string) {
	bodyData, _ := json.Marshal(types.LoginResponseType{Result: result, Code: code, Token: token})

	response.WriteHeader(200)
	response.Write(bodyData)
}

func SetLoginErrorResponse(response http.ResponseWriter, result bool, code string, message string) {
	bodyData, _ := json.Marshal(types.LoginErrorResponse{Result: result, Code: code, Message: message})

	response.WriteHeader(200)
	response.Write(bodyData)
}