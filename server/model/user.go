package model

import "gorm.io/gorm"

// User mewakili model pengguna dalam database
type User struct {
	ID        uint           `json:"id"`                      // ID adalah primary key
	Name      string         `json:"name"`                    // Name adalah nama pengguna
	Title     string         `json:"title"`                   // Title adalah judul pengguna
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"` // DeletedAt digunakan untuk soft delete
}
