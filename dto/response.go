package dto

import (
	"net/http"

	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/gin-gonic/gin"
)

func SetResponse(resCode int, data any, ctx *gin.Context) {
	response := types.ResponseDto{
		ResCode: resCode,
		DataRes: data,
	}

	ctx.JSON(http.StatusOK, response)
}