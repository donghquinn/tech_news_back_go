package middlewares

import (
	"net/http"
	"os"

	dto "github.com/dongquinn/tech_news_back_go/dto/middleware"
)

func GlobalMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authKey := os.Getenv("AUTH_KEY")

		headerKey := r.Header.Get("key")

		if headerKey == "" {
			dto.SetMiddlewareErrorResponse(w, false, "00", "No AuthKey Found")
			return
		}

		if authKey != headerKey {
			dto.SetMiddlewareErrorResponse(w, false, "00", "AuthKey is Not Valid")
			return
		}

		next.ServeHTTP(w, r)
	})
}