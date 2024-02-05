package db

import (
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	MigrateAll(db)

	fmt.Println("db migrated")

	return db
}
