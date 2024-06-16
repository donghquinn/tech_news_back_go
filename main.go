package main

import (
	"github.com/dongquinn/tech_news_back_go/configs"
	"github.com/dongquinn/tech_news_back_go/libraries/network"
	"github.com/dongquinn/tech_news_back_go/utilities"
)

func main(){
	configs.SetGlobalConfig()
	configs.SetDatabaseConfig()
	configs.SetRedisConfig()

	server := network.SetNetwork()

	utilities.GracefulShutDown(server)
}