package model

// UserValidation digunakan untuk validasi input user
type UserValidation struct {
	Name  string `json:"name"`  // Name adalah nama pengguna
	Title string `json:"title"` // Title adalah judul pengguna
}
