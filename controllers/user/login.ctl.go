package controllers

import (
	"net/http"

	"github.com/dongquinn/tech_news_back_go/auth"
	dto "github.com/dongquinn/tech_news_back_go/dto/user"
	"github.com/dongquinn/tech_news_back_go/libraries/user"
	types "github.com/dongquinn/tech_news_back_go/types/user"
	"github.com/dongquinn/tech_news_back_go/utilities"
)

func LoginController(response http.ResponseWriter, request *http.Request) {
	var loginRequest types.LoginRequestStruct

	parseErr := utilities.ParseBody(request, &loginRequest)

	if parseErr != nil {
		dto.SetLoginErrorResponse(response, false, "01", "Invalid Request")
		return
	}

	userData, userErr := user.GetLoginUserData(loginRequest.Email)

	if userErr != nil {
		dto.SetLoginErrorResponse(response, false, "02", "Get Login Data Error")
		return
	}

	token, tokenErr := auth.CreateJwtToken(userData.Uuid, userData.Email, "1")

	if tokenErr != nil {
		dto.SetLoginErrorResponse(response, false, "03", "Create Token Error")
		return
	}

	dto.SetLoginResponse(response, true, "01", token)
}