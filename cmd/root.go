package cmd

import (
	"fmt"
	"os"
)

func Run() bool {

	// Tidak ada command
	if len(os.Args) == 1 {
		return false
	}

	command := os.Args[1]

	switch command {

	case "make:model":
		MakeModel()

	case "make:controller":
		MakeController()

	case "make:migration":
		MakeMigration()

	case "make:crud":
		MakeCrud()

	case "make:middleware":
		MakeMiddleware()

	case "make:request":
		MakeRequest()

	case "migrate":
		Migrate()

	case "db:seed":
		Seed()

	case "route:list":
		RouteList()

	case "serve":
		return false

	default:
		fmt.Printf("Unknown command: %s\n", command)
		PrintHelp()
	}

	return true
}

func PrintHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  serve")
	fmt.Println("  make:model Name")
	fmt.Println("  make:controller Name")
	fmt.Println("  make:middleware Name")
	fmt.Println("  make:request Name")
	fmt.Println("  make:migration name")
	fmt.Println("  make:crud Name")
	fmt.Println("  migrate")
	fmt.Println("  db:seed")
	fmt.Println("  route:list")
}
