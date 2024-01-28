package dto

import (
	"net/http"

	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/gin-gonic/gin"
)

func SetResponse(resCode int, data any, ctx *gin.Context) {
	result := types.ResponseResult {
		Result: data,
	}

	response := types.ResponseDto{
		ResCode: resCode,
		DataRes: result,
	}

	ctx.JSON(http.StatusOK, response)
}