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
)

func main(){
	godotenv.Load(".env")
	port := os.Getenv("APP_PORT")
	mode := os.Getenv("ENVIRONMENT")

	router := gin.Default()

	gin.SetMode(mode)
	router.Use(middlewares.GlobalMiddleware())

	server := &http.Server{
		Addr: port,
		Handler: router,
	}

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Printf("Server Listening On %s\n", port)
	fmt.Printf("Server Mode: %s\n", mode)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")

	handlers.Handler(router)
	utilities.GracefulShutDown(server)
}