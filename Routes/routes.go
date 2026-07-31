package routes

import (
	controllers "golang-backend/Controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB) {
	router.GET("/health", controllers.Health)

	users := controllers.NewUserController(db)
	api := router.Group("/api")
	api.GET("/users", users.Index)
	api.POST("/users", users.Store)
}
