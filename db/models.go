package db

import "gorm.io/gorm"

type Game struct {
	ID          string `gorm:"primarykey;default:(hex(randomblob(8)))"`
	Title       string `gorm:"notnull"`
	Description string `gorm:"notnull"`
	Genre       string `gorm:"notnull"`
	ReleaseYear int    `gorm:"notnull"`
	ImageURL    string `gorm:"notnull"`
	Score       int    `gorm:"notnull"`
}

type GameSection struct {
	ID          string `gorm:"primarykey;default:(hex(randomblob(8)))"`
	Title       string `gorm:"notnull"`
	Description string `gorm:"notnull"`
	OrderOnPage int    `gorm:"notnull"`
	Games       []Game `gorm:"many2many:game_section_games"`
}

type GameSectionGames struct {
	GameSectionID  string `gorm:"primarykey"`
	GameID         string `gorm:"primarykey"`
	OrderInSection int    `gorm:"notnull"`
}

type Movie struct {
	ID          string `gorm:"primarykey;default:(hex(randomblob(8)))"`
	Title       string `gorm:"notnull"`
	Description string `gorm:"notnull"`
	Genre       string `gorm:"notnull"`
	ReleaseYear int    `gorm:"notnull"`
	ImageURL    string `gorm:"notnull"`
	Score       int    `gorm:"notnull"`
}

func MigrateAll(db *gorm.DB) {
	db.SetupJoinTable(&GameSection{}, "Games", &GameSectionGames{})

	db.AutoMigrate(&Game{}, &GameSection{}, &GameSectionGames{})
}
