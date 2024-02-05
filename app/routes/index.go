package routes

import (
	"gopplista/app/routes/admin"
	"gopplista/app/routes/games"
	"gopplista/app/routes/movies"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {

	admin.RegisterAdminRoutes(router.Group("/admin"), db)

	games.RegisterGameRoutes(router.Group("/games"), db)
	movies.RegisterMovieRoutes(router.Group("/movies"), db)

	router.Get("/", func(c *fiber.Ctx) error {

		return c.Redirect("/games")

		// return c.Render("routes/index", fiber.Map{
		// 	"Name": "Jacob",
		// }, "layouts/main", "layouts/base")
	})
}
