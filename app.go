package main

import (
	config "golang-backend/Config"
	routes "golang-backend/Routes"
	"golang-backend/cmd"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("load environment: %v", err)
	}

	if cmd.Run() {
		return
	}

	gin.SetMode(config.GinMode())
	r := gin.Default()

	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}
	routes.Register(r, config.DB)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
