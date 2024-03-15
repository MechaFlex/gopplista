package routes

import (
	"gopplista/db"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, db db.Database) {

	RegisterAdminRoutes(router.Group("/admin"), db)
	RegisterGameRoutes(router.Group("/games"), db)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/games")
	})
}
