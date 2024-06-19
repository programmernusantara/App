package database

import (
	"fmt"
	"log"

	"github.com/programmernusantara/server/model"
)

// Migrate melakukan migrasi untuk semua model yang dibutuhkan
func Migrate() {
	// Melakukan migrasi untuk model User
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Error auto migrating User model: %v", err)
	}

	fmt.Println("Database migrated successfully")
}
