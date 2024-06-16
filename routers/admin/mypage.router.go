package routers

import (
	"net/http"

	controllers "github.com/dongquinn/tech_news_back_go/controllers/admin/user"
)

func MypageRouter(server *http.ServeMux) {
	server.HandleFunc("POST /users/mypage", controllers.MyPageController)
}