package main

import (
	_ "embed"
	"gopplista/app/routes"
	"gopplista/db"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/html/v2"
	_ "github.com/joho/godotenv/autoload"
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

	envIsDev := strings.HasPrefix(os.Getenv("ENV"), "dev")
	if envIsDev {
		log.Println("Environment is set to development")
	}

	f.Use(func(c *fiber.Ctx) error {
		c.Bind(fiber.Map{
			"Development": envIsDev,
		})
		return c.Next()
	})

	routes.RegisterRoutes(f.Group("/"), database)

	f.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("routes/404", fiber.Map{
			"PageTitle": "404 | Jacob topplista",
		}, "layouts/base")
	})

	log.Fatal(f.Listen("0.0.0.0:3333"))
}
