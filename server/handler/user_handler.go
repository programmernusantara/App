package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/programmernusantara/server/database"
	"github.com/programmernusantara/server/model"
)

// UserHandler menangani permintaan ke endpoint root "/api"
// Endpoint ini mengembalikan pesan sambutan
func UserHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to the User API")
}

// GetUser menangani permintaan untuk mendapatkan semua user
// Endpoint ini mengembalikan daftar semua user dalam bentuk JSON
func GetUser(c *fiber.Ctx) error {
	var users []model.User
	// Mengambil semua user dari database
	database.DB.Find(&users)
	// Mengembalikan data user dalam format JSON
	return c.JSON(users)
}

// UserGetById menangani permintaan untuk mendapatkan user berdasarkan ID
// Endpoint ini mengembalikan data user yang sesuai dengan ID yang diberikan
func UserGetById(c *fiber.Ctx) error {
	id := c.Params("id") // Mengambil ID dari parameter URL
	var user model.User
	// Mencari user berdasarkan ID
	if result := database.DB.First(&user, id); result.Error != nil {
		// Jika user tidak ditemukan, kembalikan status 404
		return c.Status(404).SendString("User not found")
	}
	// Mengembalikan data user dalam format JSON
	return c.JSON(user)
}

// PostUser menangani permintaan untuk menambahkan user baru
// Endpoint ini menerima data user dalam bentuk JSON dan menambahkannya ke database
func PostUser(c *fiber.Ctx) error {
	var data model.UserValidation
	// Mem-parsing data dari request body ke struct UserValidation
	if err := c.BodyParser(&data); err != nil {
		// Jika ada kesalahan saat parsing, kembalikan status 400
		return c.Status(400).JSON(err.Error())
	}
	// Membuat instance User baru dengan data yang diterima
	user := model.User{
		Name:  data.Name,
		Title: data.Title,
	}
	// Menyimpan user baru ke database
	database.DB.Create(&user)
	// Mengembalikan data user yang baru dibuat dalam format JSON
	return c.JSON(user)
}

// UpdateUser menangani permintaan untuk mengupdate user berdasarkan ID
// Endpoint ini menerima data user dalam bentuk JSON dan mengupdate data user yang sesuai dengan ID yang diberikan
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id") // Mengambil ID dari parameter URL
	var data model.UserValidation
	// Mem-parsing data dari request body ke struct UserValidation
	if err := c.BodyParser(&data); err != nil {
		// Jika ada kesalahan saat parsing, kembalikan status 400
		return c.Status(400).JSON(err.Error())
	}

	var user model.User
	// Mencari user berdasarkan ID
	if result := database.DB.First(&user, id); result.Error != nil {
		// Jika user tidak ditemukan, kembalikan status 404
		return c.Status(404).SendString("User not found")
	}

	// Mengupdate field Name dan Title dari user yang ditemukan
	user.Name = data.Name
	user.Title = data.Title
	// Menyimpan perubahan ke database
	database.DB.Save(&user)
	// Mengembalikan data user yang telah diperbarui dalam format JSON
	return c.JSON(user)
}

// DeleteUser menangani permintaan untuk menghapus user berdasarkan ID
// Endpoint ini menghapus data user yang sesuai dengan ID yang diberikan
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id") // Mengambil ID dari parameter URL
	var user model.User
	// Mencari user berdasarkan ID
	if result := database.DB.First(&user, id); result.Error != nil {
		// Jika user tidak ditemukan, kembalikan status 404
		return c.Status(404).SendString("User not found")
	}
	// Menghapus user dari database
	database.DB.Delete(&user)
	// Mengembalikan pesan sukses
	return c.SendString("User deleted successfully")
}
