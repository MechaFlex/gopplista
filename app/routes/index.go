package routes

import (
	"gopplista/app/routes/games"
	"gopplista/db"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, db db.Database) {

	//admin.RegisterAdminRoutes(router.Group("/admin"), db)

	games.RegisterGameRoutes(router.Group("/games"), db)
	//movies.RegisterMovieRoutes(router.Group("/movies"), db)

	router.Get("/", func(c *fiber.Ctx) error {

		return c.Redirect("/games")

		// return c.Render("routes/index", fiber.Map{
		// 	"Name": "Jacob",
		// }, "layouts/main", "layouts/base")
	})
}
