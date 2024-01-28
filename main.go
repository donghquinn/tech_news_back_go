package main

import (
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
	router.Use(middlewares.CorsMiddlewares())

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