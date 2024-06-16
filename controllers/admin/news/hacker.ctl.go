package controllers

import (
	"net/http"

	"github.com/dongquinn/tech_news_back_go/auth"
	dto "github.com/dongquinn/tech_news_back_go/dto/news"
	"github.com/dongquinn/tech_news_back_go/libraries/news"
	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/dongquinn/tech_news_back_go/utilities"
)

func AdminHackerNewsController(response http.ResponseWriter, request *http.Request){
	_, _, _, err := auth.ValidateJwtToken(request)

	if err != nil {
		dto.SetHackerErrorResponse(response, false, "01", "JWT Verifying Error")
		return
	}

	requestBody := types.HackerNewsRequest{}

	parseErr := utilities.ParseBody(request, &requestBody)

	if parseErr != nil {
		dto.SetHackerErrorResponse(response, false, "01", "Not Valid Request")
		return
	}

	page := request.URL.Query().Get("page")
	size := request.URL.Query().Get("size")

	newsList, newsErr := news.GetTodayHackerNewsList(page, size)

	if newsErr != nil {
		dto.SetHackerErrorResponse(response, false, "02", "Query News Error")
		return
	}

	dto.SetHackerResponse(response, true, "01", newsList)
}