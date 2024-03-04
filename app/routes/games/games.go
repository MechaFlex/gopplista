package games

import (
	"fmt"
	"gopplista/db"

	"github.com/gofiber/fiber/v2"
)

func RegisterGameRoutes(router fiber.Router, db db.Database) {

	router.Get("/", func(c *fiber.Ctx) error {

		gameSectionsWithGames, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("routes/games/games", fiber.Map{
			"PageTitle": "Spel | Jacobs topplista",
			"Sections":  gameSectionsWithGames,
		}, "layouts/main", "layouts/base")
	})
}
