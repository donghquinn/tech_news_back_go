package network

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dongquinn/tech_news_back_go/configs"
	"github.com/dongquinn/tech_news_back_go/middlewares"
	"github.com/dongquinn/tech_news_back_go/routers"
	admin "github.com/dongquinn/tech_news_back_go/routers/admin"
)


func SetNetwork() *http.Server {
	globalConfig := configs.GlobalConfig

	handler := http.NewServeMux()

	corsHandler := middlewares.CorsMiddlewares(handler)
	middlewareHandler := middlewares.GlobalMiddleware(corsHandler)
	
	routers.GeekRouter(handler)
	routers.HackerRouter(handler)
	routers.MlRouter(handler)

	admin.AdminGeekRouter(handler)
	admin.AdminHackerRouter(handler)
	admin.AdminMlRouter(handler)
	admin.MypageRouter(handler)

	server := &http.Server {
		Handler: middlewareHandler,
		Addr: globalConfig.AppPort,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
	}

	startTime := time.Now()

	message := fmt.Sprintf("Server Start At : %s", startTime)
	message2 :=  fmt.Sprintf("Listening ON : %s", globalConfig.AppPort)
	wrapper := strings.Repeat("@", len(message))

	log.Printf("%s", wrapper)
	log.Printf("%s", message)
	log.Printf("%s", message2)
	log.Printf("%s", wrapper)

	return server
}