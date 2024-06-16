package dto

import (
	"encoding/json"
	"net/http"

	types "github.com/dongquinn/tech_news_back_go/types/user"
)

func SetSignupResponse(response http.ResponseWriter, result bool, code string) {
	responseBody, _ := json.Marshal(types.SignupResponse{Result: result, Code: code})

	response.WriteHeader(200)
	response.Write(responseBody)
}

func SetSignupErroResponse(response http.ResponseWriter, result bool, code string, message string) {
	responseBody, _ := json.Marshal(types.SignupErrorResponse{Result: result, Code: code, Message: message})

	response.WriteHeader(200)
	response.Write(responseBody)
}