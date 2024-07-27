package controllers

import (
	"net/http"

	dto "github.com/dongquinn/tech_news_back_go/dto/news"
	"github.com/dongquinn/tech_news_back_go/libraries/news"
	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/dongquinn/tech_news_back_go/utilities"
)

func GeekNewsController(response http.ResponseWriter, request *http.Request) {
	requestBody := types.GeekNewsRequest{}

	parseErr := utilities.ParseBody(request, &requestBody)

	if parseErr != nil {
		dto.SetGeekErrorResponse(response, false, "01", "Not Valid Request")
		return
	}

	page := request.URL.Query().Get("page")
	size := request.URL.Query().Get("size")


	geekNewsList, geekErr := news.GetTodayGeekNewsList(requestBody.Today, page, size)

	if geekErr != nil {
		dto.SetGeekErrorResponse(response, false, "02", "Query Geek News Error")
		return
	}

	dto.SetGeekResponse(response, true, "01", geekNewsList)
}