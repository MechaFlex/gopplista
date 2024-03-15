package main

import (
	"embed"
	"gopplista/app/routes"
	"gopplista/db"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/html/v2"
	_ "github.com/joho/godotenv/autoload"
)

//go:embed app/*
var rootAppFS embed.FS

func main() {
	// Initialize database
	database, err := db.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Initialize Fiber app with templates
	appFS, err := fs.Sub(rootAppFS, "app/templates")
	if err != nil {
		log.Fatal(err)
	}
	f := fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(appFS), ".html"),
	})

	// Things that should happen before every request
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

	f.Use(logger.New())

	// Register routes, main part of the app
	routes.RegisterRoutes(f.Group("/"), database)

	// Serve static files
	staticFS, _ := fs.Sub(rootAppFS, "app/static")
	f.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(staticFS),
	}))

	// If no route matches this is the catch all fallback
	f.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("pages/404", fiber.Map{
			"PageTitle": "404 | Jacob topplista",
		}, "layouts/base")
	})

	log.Fatal(f.Listen("0.0.0.0:3333"))
}
