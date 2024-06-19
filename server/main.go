package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/programmernusantara/server/database"
	"github.com/programmernusantara/server/router"
)

func main() {
	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Mengaktifkan CORS agar aplikasi dapat menerima permintaan dari berbagai domain
	app.Use(cors.New())

	// Koneksi ke database dan melakukan migrasi
	database.Connect()
	database.Migrate()

	// Mengatur routing untuk aplikasi
	router.UserRouter(app)

	// Menjalankan server pada port 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
