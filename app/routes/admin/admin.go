package admin

import (
	"fmt"
	games "gopplista/app/routes/admin/games"
	dbpkg "gopplista/db"
	"os"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterAdminRoutes(router fiber.Router, db dbpkg.Database) {

	pasetoKey := paseto.NewV4SymmetricKey()

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		log.Warn("The environment variable ADMIN_PASSWORD is not set. Admin password is an empty string.")
	}

	router.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/admin/login" {
			return c.Next()
		}

		if c.Cookies("admin") == "" {
			return c.Redirect("/admin/login")
		}

		tokenString := c.Cookies("admin")
		parser := paseto.NewParser()
		_, err := parser.ParseV4Local(pasetoKey, tokenString, nil)
		if err != nil {
			return c.Redirect("/admin/login")
		}

		return c.Next()
	})

	games.RegisterGameSectionRoutes(router.Group("/games/sections"), db) //Put before to prioritize specifics over dynamics
	games.RegisterGamesGamesRoutes(router.Group("/games"), db)

	router.Get("/", func(c *fiber.Ctx) error {

		gameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		allGames, err := db.Queries.GetGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting games: %v", err))
		}

		return c.Render("routes/admin/index", fiber.Map{
			"Sections": gameSections,
			"Games":    allGames,
		}, "layouts/admin", "layouts/base")
	})

	router.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("routes/admin/login", nil, "layouts/base")
	})

	router.Post("/login", func(c *fiber.Ctx) error {

		if c.FormValue("password") != adminPassword {
			return c.Status(400).SendString("Wrong password")
		}

		token := paseto.NewToken()
		token.SetIssuedAt(time.Now())
		token.SetExpiration(time.Now().Add(24 * time.Hour))
		encryptedToken := token.V4Encrypt(pasetoKey, nil)

		c.Cookie(&fiber.Cookie{
			Name:     "admin",
			Value:    encryptedToken,
			Secure:   true,
			HTTPOnly: true,
			SameSite: "Strict",
		})

		return c.Status(200).SendString("Logged in")
	})
}
