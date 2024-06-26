package routers

import (
	"net/http"

	controllers "github.com/dongquinn/tech_news_back_go/controllers/admin/news"
)

func AdminMlRouter(server *http.ServeMux) {
	server.HandleFunc("POST /admin/ml/latest", controllers.AdminMachineLearningController)
}