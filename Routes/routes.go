package Routes

import (
	"golang-backend/Controller"

	"github.com/gin-gonic/gin"
)

func RoutesInit(router *gin.Engine) {
	router.GET("/", Controller.Tes)
	router.GET("/bismillah-tes", Controller.GetUser)
	router.GET("/bismillah-create", Controller.CreateUser)
}