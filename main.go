package main

import (
	_ "embed"
	"gopplista/app/routes"
	"gopplista/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database, err := db.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	f := fiber.New(fiber.Config{
		Views: html.New("./app", ".html"),
	})

	f.Use(logger.New())

	f.Static("/", "./app/static")

	app := f.Group("/")
	routes.RegisterRoutes(app, database)

	f.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("routes/404", nil, "layouts/base")
	})

	log.Fatal(f.Listen("localhost:3333"))
}
