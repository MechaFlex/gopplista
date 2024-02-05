package movies

import (
	dbpkg "gopplista/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterMovieRoutes(router fiber.Router, db *gorm.DB) {

	var movies []dbpkg.Movie
	db.Find(&movies)
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("routes/movies/movies", fiber.Map{
			"Movies": movies,
		}, "layouts/main", "layouts/base")
	})
}
