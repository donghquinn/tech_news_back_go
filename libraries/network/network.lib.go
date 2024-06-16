package network

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dongquinn/tech_news_back_go/configs"
	"github.com/dongquinn/tech_news_back_go/routers"
)


func SetNetwork() *http.Server {
	globalConfig := configs.GlobalConfig

	handler := http.NewServeMux()

	routers.GeekRouter(handler)
	routers.HackerRouter(handler)
	routers.MlRouter(handler)

	server := &http.Server {
		Handler: handler,
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