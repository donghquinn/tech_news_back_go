package controllers

import (
	"net/http"

	dto "github.com/dongquinn/tech_news_back_go/dto/news"
	"github.com/dongquinn/tech_news_back_go/libraries/news"
	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/dongquinn/tech_news_back_go/utilities"
)

func MachineLearningController(response http.ResponseWriter, request *http.Request) {
	requestBody := types.MachineLearningNewsRequest{}

	parseErr := utilities.ParseBody(request, &requestBody)

	if parseErr != nil {
		dto.SetMlErrorResponse(response, false, "01", "Not Valid Error")
		return
	}
	
	page := request.URL.Query().Get("page")
	size := request.URL.Query().Get("size")

	mlNewsList, mlErr := news.GetTodayMlNewsList(requestBody.Today, page, size)

	if mlErr != nil {
		dto.SetMlErrorResponse(response, false, "02", "Get ML News Error")
		return
	}
	
	dto.SetMlResponse(response, true, "01", mlNewsList)
}