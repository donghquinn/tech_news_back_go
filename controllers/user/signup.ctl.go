package controllers

import (
	"net/http"

	dto "github.com/dongquinn/tech_news_back_go/dto/user"
	"github.com/dongquinn/tech_news_back_go/libraries/user"
	types "github.com/dongquinn/tech_news_back_go/types/user"
	"github.com/dongquinn/tech_news_back_go/utilities"
)

func SignupController(response http.ResponseWriter, request *http.Request) {
	var requestBody types.SignupRequestStruct

	parseErr := utilities.ParseBody(request, &requestBody)

	if parseErr != nil {
		dto.SetSignupErroResponse(response, false, "01", "Invalid Request Error")
		return
	}

	signupErr := user.SignupUser(requestBody.Email, requestBody.Name, requestBody.Password)

	if signupErr != nil {
		dto.SetSignupErroResponse(response, false, "02", "Signup Error")
		return
	}

	dto.SetSignupResponse(response, true, "01")
}