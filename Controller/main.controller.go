package Controller

import (
	"golang-backend/Config"
	"golang-backend/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Tes(c *gin.Context){
	c.String(http.StatusOK, "Bismillah test")
}

func CreateUser(c *gin.Context){
	var createUser []Model.User
	c.BindJSON(&createUser)
	Config.DB.Create(&createUser)
	
}

func GetUser(c *gin.Context){
	var getUser []Model.User
	Config.DB.Find(&getUser)
	c.JSON(http.StatusOK, getUser)
}