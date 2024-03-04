package routes

import (
	"gopplista/app/routes/admin"
	"gopplista/app/routes/games"
	"gopplista/db"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, db db.Database) {

	admin.RegisterAdminRoutes(router.Group("/admin"), db)

	games.RegisterGameRoutes(router.Group("/games"), db)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/games")
	})
}
