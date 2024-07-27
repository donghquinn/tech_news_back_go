package routers

import (
	"net/http"

	controllers "github.com/dongquinn/tech_news_back_go/controllers/news"
)

func GeekRouter(server *http.ServeMux) {
	server.HandleFunc("POST /geek/news", controllers.GeekNewsController)
}