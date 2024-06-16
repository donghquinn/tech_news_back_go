package controllers

import (
	"net/http"

	dto "github.com/dongquinn/tech_news_back_go/dto/user"
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

	
}