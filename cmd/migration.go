package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func MakeMigration() {
	if len(os.Args) < 3 || strings.TrimSpace(os.Args[2]) == "" {
		fmt.Println("Usage: make:migration name")
		return
	}

	name := snakeCase(strings.ReplaceAll(strings.TrimSpace(os.Args[2]), "-", "_"))
	for _, r := range name {
		if !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_') {
			fmt.Println("Migration name may only contain letters, numbers, hyphens, and underscores.")
			return
		}
	}

	filename := filepath.Join("Migrations", time.Now().Format("20060102150405")+"_"+name+".up.sql")
	content := fmt.Sprintf("-- Migration: %s\n-- Write the SQL needed to apply this change.\n\n", name)
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		fmt.Printf("Create directory: %v\n", err)
		return
	}
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		fmt.Printf("Create migration: %v\n", err)
		return
	}
	fmt.Printf("Created: %s\n", filename)
}
