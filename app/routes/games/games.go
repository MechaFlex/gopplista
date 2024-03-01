package games

import (
	"fmt"
	dbpkg "gopplista/db"
	models "gopplista/db/gen"

	"github.com/gofiber/fiber/v2"
)

type GameSectionWithGames struct {
	models.GameSection
	Games []models.Game
}

func RegisterGameRoutes(router fiber.Router, db dbpkg.Database) {

	router.Get("/", func(c *fiber.Ctx) error {

		sections, err := db.Queries.GetGameSections(db.Ctx)
		if err != nil {
			return err
		}

		sectionsWithGames := []GameSectionWithGames{}

		for _, section := range sections {
			games, err := db.Queries.GetGamesInGameSection(db.Ctx, section.ID)
			if err != nil {
				return err
			}
			sectionsWithGames = append(sectionsWithGames, GameSectionWithGames{section, games})
		}

		fmt.Println(sectionsWithGames)

		fmt.Println("games:", sectionsWithGames[0].Games)

		return c.Render("routes/games/games", fiber.Map{
			"Sections": sectionsWithGames,
		}, "layouts/main", "layouts/base")
	})
}
