package main

import (
	config "golang-backend/Config"
	routes "golang-backend/Routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	err := godotenv.Load()
	if err != nil {
		panic("Gagal Load ENV")
	}
	
	r := gin.Default()

	config.ConnectionDB()
	routes.RoutesInit(r)

	r.Run(os.Getenv("APP_PORT"))
}