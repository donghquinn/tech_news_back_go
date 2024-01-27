package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dongquinn/tech_news_back_go/handlers"
	"github.com/dongquinn/tech_news_back_go/middlewares"
	"github.com/dongquinn/tech_news_back_go/utilities"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main(){
	godotenv.Load(".env")
	port := os.Getenv("APP_PORT")
	mode := os.Getenv("ENVIRONMENT")

	router := gin.Default()

	gin.SetMode(mode)

	router.Use(middlewares.GlobalMiddleware())
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(400)
		} else {
			c.Next()
		}
	})
	server := &http.Server{
		Addr: port,
		Handler: router,
	}

	log := utilities.Logger()

	log.Info("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	log.Info("Server Listening On", zap.String("port",port))
	log.Info("Server Mode: ", zap.String("mode", mode) )
	log.Info("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

	handlers.Handler(router)
	utilities.GracefulShutDown(server)
}