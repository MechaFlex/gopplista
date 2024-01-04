package main

import (
	"fmt"
	"log"

	dbpkg "gopplista/db"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db/db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dbpkg.MigrateAll(db)

	fmt.Println("db migrated")

	engine := html.New("./views", ".html")

	F := fiber.New(fiber.Config{
		Views: engine,
	})

	F.Use(logger.New())

	F.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	F.Get("/teststatic", func(c *fiber.Ctx) error {
		return c.SendFile("./static.html")
	})

	var games []dbpkg.Game
	db.Find(&games)
	fmt.Println(games)
	F.Get("/base", func(c *fiber.Ctx) error {
		return c.Render("base", fiber.Map{
			"Name":  "Jacob",
			"Games": games,
		})
	})

	log.Fatal(F.Listen(":3333"))
}
