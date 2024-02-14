package newscontrollers

import (
	"log"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/dto"
	"github.com/dongquinn/tech_news_back_go/libraries/geek"
	"github.com/dongquinn/tech_news_back_go/libraries/prisma"
	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/gin-gonic/gin"
)

func GeekNewsController(ctx *gin.Context) {
	request := types.GeekNewsRequest{}

	if reqErr := ctx.ShouldBind(&request); reqErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required Value is not included. Please try again."})

		return
	}

	today := request.Today

	client, dbErr := prisma.PrismaClient()

	if dbErr != nil {
		log.Fatalln(dbErr)
	}

	result := geek.GetDailyGeekNews(client, today)

	dto.SetResponse(200, result, ctx)
}