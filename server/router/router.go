package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/programmernusantara/server/handler"
)

// UserRouter mengatur routing untuk endpoint user
func UserRouter(app *fiber.App) {
	api := app.Group("/api") // Membuat grup route "/api"

	// Mendefinisikan route untuk user
	api.Get("/", handler.UserHandler)           // Route untuk mengakses endpoint root "/api"
	api.Get("/user", handler.GetUser)           // Route untuk mendapatkan semua user
	api.Get("/user/:id", handler.UserGetById)   // Route untuk mendapatkan user berdasarkan ID
	api.Post("/user", handler.PostUser)         // Route untuk menambahkan user baru
	api.Put("/user/:id", handler.UpdateUser)    // Route untuk mengupdate user berdasarkan ID
	api.Delete("/user/:id", handler.DeleteUser) // Route untuk menghapus user berdasarkan ID
}
