package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GlobalMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		start := time.Now()
		
		secretKey := os.Getenv("AUTH_KEY")

		log.Printf("Received Request: %s", start.String())

		headerKey := ctx.Request.Header.Get("key")

		log.Printf("Check Header Key: %s.\nStart Validate with Secret Key", headerKey)

		// Header Key Matching
		if headerKey != secretKey {
			ctx.JSON(
				http.StatusBadRequest, 
				gin.H{"message": "Header Key is Not Valid. Please Check and Try Again."})
		}

		ctx.Next()

		latency := time.Since(start)

		log.Printf("Latency: %s", latency.String())
		ctx.Writer.Status()
	}
}