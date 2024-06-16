package routers

import (
	"net/http"

	controllers "github.com/dongquinn/tech_news_back_go/controllers/user"
)

func UserRouter(server *http.ServeMux) {
	server.HandleFunc("POST /users/login", controllers.LoginController)
}