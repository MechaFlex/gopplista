package db

import "gorm.io/gorm"

type Game struct {
	ID          string `gorm:"primarykey;default:(hex(randomblob(8)));"`
	Title       string `gorm:"notnull"`
	Description string
	Genre       string
	ReleaseYear uint
	ImageURL    string `gorm:"notnull"`
	Score       uint   `gorm:"notnull"`
}

func MigrateAll(db *gorm.DB) {
	db.AutoMigrate(&Game{})
}
