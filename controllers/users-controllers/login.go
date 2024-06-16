package userscontrollers

import (
	"log"
	"net/http"

	"github.com/dongquinn/tech_news_back_go/dto"
	"github.com/dongquinn/tech_news_back_go/libraries/users"
	"github.com/dongquinn/tech_news_back_go/types"
	"github.com/gin-gonic/gin"
)

// TODO Redis로 로그인 중복 요청 체크
func UserLoginController(ctx *gin.Context) {
	request := types.LoginRequest{}

	if reqErr := ctx.ShouldBind(&request); reqErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Request Is Not Valid. Please Try Again."})

		return
	}

	email := request.Email
	password := request.Password

	log.Printf("Login Email: %s, Password: %s\n", email, password)

	result, loginErr := users.LoginLibrary(email, password)

	if loginErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": loginErr})

		return
	}

	dto.SetResponse(200, result, ctx)
}