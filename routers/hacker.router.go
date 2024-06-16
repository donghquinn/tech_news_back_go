package routers

import (
	"net/http"

	controllers "github.com/dongquinn/tech_news_back_go/controllers/news"
)

func HackerRouter(server *http.ServeMux) {
	server.HandleFunc("POST /hacker/news", controllers.HackerNewsController)

}