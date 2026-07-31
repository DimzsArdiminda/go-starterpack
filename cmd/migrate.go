package cmd

import (
	"errors"
	"fmt"
	config "golang-backend/Config"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

type migrationRecord struct {
	Name      string    `gorm:"primaryKey;size:255"`
	AppliedAt time.Time `gorm:"not null"`
}

func (migrationRecord) TableName() string { return "schema_migrations" }

func Migrate() {
	if err := config.Connect(); err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		return
	}
	if err := config.DB.AutoMigrate(&migrationRecord{}); err != nil {
		fmt.Printf("Create migration table: %v\n", err)
		return
	}

	files, err := filepath.Glob(filepath.Join("Migrations", "*.up.sql"))
	if err != nil {
		fmt.Printf("Find migrations: %v\n", err)
		return
	}
	sort.Strings(files)
	if len(files) == 0 {
		fmt.Println("No migration files found.")
		return
	}

	applied := 0
	for _, file := range files {
		name := filepath.Base(file)
		var record migrationRecord
		err := config.DB.First(&record, "name = ?", name).Error
		if err == nil {
			continue
		}
		if !isNotFound(err) {
			fmt.Printf("Read migration history for %s: %v\n", name, err)
			return
		}
		contents, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Read %s: %v\n", file, err)
			return
		}
		if err := applyMigration(name, string(contents)); err != nil {
			fmt.Printf("Apply %s: %v\n", name, err)
			return
		}
		fmt.Printf("Migrated: %s\n", name)
		applied++
	}
	fmt.Printf("Migration complete (%d applied).\n", applied)
}

func applyMigration(name, source string) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		for _, statement := range strings.Split(source, ";") {
			statement = strings.TrimSpace(statement)
			if statement == "" || strings.HasPrefix(statement, "--") {
				continue
			}
			if err := tx.Exec(statement).Error; err != nil {
				return err
			}
		}
		return tx.Create(&migrationRecord{Name: name, AppliedAt: time.Now()}).Error
	})
}

func isNotFound(err error) bool { return errors.Is(err, gorm.ErrRecordNotFound) }
