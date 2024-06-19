package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database
var DB *gorm.DB

// Connect menghubungkan aplikasi ke database PostgreSQL
func Connect() {
	// Memuat konfigurasi dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Membuat string koneksi database dari konfigurasi .env
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Membuka koneksi ke database
	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Failed to connect to database")
	}

	fmt.Println("Database connected successfully")
}
