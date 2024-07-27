package controllers

import (
	"net/http"

	"github.com/dongquinn/tech_news_back_go/auth"
	dto "github.com/dongquinn/tech_news_back_go/dto/user"
	"github.com/dongquinn/tech_news_back_go/libraries/user"
)

// TODO  라우터 분리 - 뉴스별로
func MyPageController(response http.ResponseWriter, request *http.Request) {
	userId, email, _, err := auth.ValidateJwtToken(request)

	if err != nil {
		dto.SetMypageErrorResponse(response, false, "01", "JWT Verifying Error")
		return
	}

	page := request.URL.Query().Get("page")
	size := request.URL.Query().Get("size")

	hackerNewsList, mlNewsList, geekNewsList, mypageErr := user.GetMyPage(userId, email, page, size)

	if mypageErr != nil {
		dto.SetMypageErrorResponse(response, false, "02", "Get Liked News List Error")
		return
	}

	dto.SetMypageResponse(response, true, "01", hackerNewsList, mlNewsList, geekNewsList)
}