package newscontrollers

import (
	"log"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/dto"
	"github.com/dongquinn/tech_news_back_go/libraries/ml"
	"github.com/dongquinn/tech_news_back_go/libraries/prisma"
	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/gin-gonic/gin"
)

func MachineLearningController(ctx *gin.Context) {
	request := types.MachineLearningNewsRequest{}

	if reqErr := ctx.ShouldBind(&request); reqErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required Value is not included. Please Try Again"})

		return
	}
	
	today := request.Today

	page := ctx.Query("page")
	size := ctx.Query("size")

	client, prismaErr := prisma.PrismaClient()

	if prismaErr != nil {
		log.Fatalln(prismaErr)
	}

	result := ml.GetMlNews(client, today, page, size)

	dto.SetResponse(200, result, ctx)
}