package controllers

import (
	"log"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/dto"
	"github.com/dongquinn/tech_news_back_go/libraries/hacker"
	"github.com/dongquinn/tech_news_back_go/libraries/prisma"
	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/gin-gonic/gin"
)

func HackerNewsController(ctx *gin.Context){
	request := types.HackerNewsRequest{}

	if reqErr := ctx.ShouldBind(&request); reqErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required value is not included"})

		return
	}

	client, prismaErr := prisma.PrismaClient()

	if prismaErr != nil {
		log.Fatalln(prismaErr)
	}

	today := request.Today

	result := hacker.GetHackerNews(client, today)

	dto.SetResponse(200, result, ctx)
}