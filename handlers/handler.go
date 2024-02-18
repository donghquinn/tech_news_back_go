package handlers

import (
	newscontrollers "github.com/dongquinn/tech_news_back_go/controllers/news-controllers"
	userscontrollers "github.com/dongquinn/tech_news_back_go/controllers/users-controllers"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	hacker := router.Group("/hacker")
	{
		hacker.POST("/news", newscontrollers.HackerNewsController)
	}

	geek := router.Group("/geek")
	{
		geek.POST("/news", newscontrollers.GeekNewsController)
	}

	ml := router.Group("/ml")
	{
		ml.POST("/latest", newscontrollers.MachineLearningController)
	}

	login := router.Group("/users")
 
	{
		login.POST("/login", userscontrollers.UserLoginController)
	}

}