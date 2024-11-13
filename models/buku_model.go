package models

import (
	"time"

	"gorm.io/gorm"
)

// Buku is a struct that represents a book
type Buku struct {
	Id        *int      ` gorm:"primaryKey;not null" `
	Judul     string    `gorm:"not null" `
	Harga     *int      `gorm:"not null" `
	CreatedAt time.Time `gorm:"autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Migrate migrates the schema
func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&Buku{}); err != nil {
		panic("Migration failed: " + err.Error())
	}
}
