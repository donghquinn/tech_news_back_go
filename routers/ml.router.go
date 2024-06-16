package routers

import (
	"net/http"

	controllers "github.com/dongquinn/tech_news_back_go/controllers/news"
)

func MlRouter(server *http.ServeMux) {
	server.HandleFunc("POST /ml/latest", controllers.MachineLearningController)
}