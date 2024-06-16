package routers

import (
	"net/http"

	controllers "github.com/dongquinn/tech_news_back_go/controllers/admin/news"
)

func AdminGeekRouter(server *http.ServeMux) {
	server.HandleFunc("POST /geek/news", controllers.AdminGeekNewsController)
}