package handlers

import (
	"github.com/dongquinn/tech_news_back_go/controllers"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	hacker :=router.Group("/hacker")
	{
		hacker.POST("/news", controllers.HackerNewsController)
	}

	geek := router.Group("/hada")
	{
		geek.POST("/news", controllers.GeekNewsController)
	}

	ml := router.Group("/ml")
	{
		ml.POST("/latest", controllers.MachineLearningController)
	}
}