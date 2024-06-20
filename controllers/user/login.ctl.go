package controllers

import (
	"net/http"

	"github.com/dongquinn/tech_news_back_go/auth"
	dto "github.com/dongquinn/tech_news_back_go/dto/user"
	"github.com/dongquinn/tech_news_back_go/libraries/crypt"
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

	decodedEmail, decodedPassword, decodeErr := user.GetDecodeUserData(loginRequest)

	if decodeErr != nil {
		dto.SetLoginErrorResponse(response, false, "02", "Decode Request Error")
		return
	}

	userData, userErr := user.GetLoginUserData(decodedEmail)

	if userErr != nil {
		dto.SetLoginErrorResponse(response, false, "02", "Get Login Data Error")
		return
	}

	isMatch, passwordMatchErr := crypt.PasswordCompare(userData.Password, decodedPassword)

	if passwordMatchErr != nil || !isMatch {
		dto.SetLoginErrorResponse(response, false, "04", "Password Is Invalid")
		return
	}

	token, tokenErr := auth.CreateJwtToken(userData.Uuid, userData.Email, userData.UserStatus)

	if tokenErr != nil {
		dto.SetLoginErrorResponse(response, false, "03", "Create Token Error")
		return
	}

	dto.SetLoginResponse(response, true, "01", token)
}