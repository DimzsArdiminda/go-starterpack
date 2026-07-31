package cmd

import (
	"fmt"
	config "golang-backend/Config"
	"golang-backend/Database"
)

func Seed() {
	if err := config.Connect(); err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		return
	}
	if err := database.Seed(config.DB); err != nil {
		fmt.Printf("Seed failed: %v\n", err)
		return
	}
	fmt.Println("Database seed complete.")
}
