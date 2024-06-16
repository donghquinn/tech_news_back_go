package dto

import (
	"encoding/json"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/types"
	mypage "github.com/dongquinn/tech_news_back_go/types/user"
)

func SetMypageResponse(response http.ResponseWriter, result bool, code string, hackerNews []types.HackerNewsResponse, mlNews []types.MachineLEarningNewsResponse, geekNews []types.GeekNewsResponse) {
	responseBody, _ := json.Marshal(mypage.MyPageResponse{Result: result, Code: code, HackerNews: hackerNews, MlNews: mlNews, GeekNews: geekNews})

	response.WriteHeader(200)
	response.Write(responseBody)
}

func SetMypageErrorResponse(response http.ResponseWriter, result bool, code string, message string) {
	responseBody, _ := json.Marshal(mypage.MyPageErrorResponse{Result: result, Code: code, Message: message})

	response.WriteHeader(200)
	response.Write(responseBody)
}