package games

import (
	dbpkg "gopplista/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterGameRoutes(router fiber.Router, db *gorm.DB) {

	router.Get("/", func(c *fiber.Ctx) error {
		var sections []dbpkg.GameSection
		db.Preload("Games").Find(&sections)

		return c.Render("routes/games/games", fiber.Map{
			"Sections": sections,
		}, "layouts/main", "layouts/base")
	})
}
