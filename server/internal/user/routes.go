package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanmintegui/edina/internal/database"
	"github.com/nathanmintegui/edina/internal/user/dto"
)

var service Service

func Configure() {
	service = Service{
		Repository: &RepositoryPostgres{
			Conn: database.Conn,
		},
	}
}

func SetRoutes(router *gin.Engine) {
	router.POST("/login", handleLogin)
}

func handleLogin(ctx *gin.Context) {
	var requestBody dto.LoginRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "malformed JSON, invalid structure",
		})
		return
	}

	if requestBody.Login == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "login field is missing",
		})
		return
	}

	if requestBody.Password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "password field is missing",
		})
		return
	}

	token, err := service.Login(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
