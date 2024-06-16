package dto

import (
	"net/http"

	"github.com/dongquinn/tech_news_back_go/types"
)

func SetResponse(response http.ResponseWriter, result bool, code string) {
	result := types.ResponseResult {
		Result: data,
	}

	response := types.ResponseDto{
		ResCode: resCode,
		DataRes: result,
	}

	ctx.JSON(http.StatusOK, response)
}