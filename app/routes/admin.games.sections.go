package routes

import (
	"fmt"
	dbpkg "gopplista/db"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func RegisterGameSectionRoutes(router fiber.Router, db dbpkg.Database) {

	// Get game sections
	router.Get("/", func(c *fiber.Ctx) error {

		gameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("pages/admin/games/sectionslist", fiber.Map{
			"Sections": gameSections,
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {

		_, err := db.Queries.CreateGameSection(db.Ctx, dbpkg.CreateGameSectionParams{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
		})
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error creating game section: %v", err))
		}

		updatedGameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("pages/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	// Update order of game sections
	router.Put("/", func(c *fiber.Ctx) error {

		payload := struct {
			SectionIDs []string `form:"section"`
		}{}

		err := c.BodyParser(&payload)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(payload)

		for i, sectionID := range payload.SectionIDs {
			db.Queries.UpdateGameSectionOrder(db.Ctx, dbpkg.UpdateGameSectionOrderParams{
				ID:          sectionID,
				OrderOnPage: int64(i),
			})
		}

		updatedGameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("pages/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	// Update content of section
	router.Put("/:id", func(c *fiber.Ctx) error {

		_, err := db.Queries.UpdateGameSection(db.Ctx, dbpkg.UpdateGameSectionParams{
			ID:          c.Params("id"),
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
		})
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error updating game section: %v", err))
		}

		updatedGameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("pages/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	// Delete section
	router.Delete("/:id", func(c *fiber.Ctx) error {

		_, err := db.Queries.DeleteGameSection(db.Ctx, dbpkg.GameSection{
			ID: c.Params("id"),
		})
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error deleting game section: %v", err))
		}

		updatedGameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("pages/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	// Set games in section
	router.Put("/:sectionid/games", func(c *fiber.Ctx) error {

		payload := struct {
			GameIDs []string `form:"game"`
		}{}

		err := c.BodyParser(&payload)
		if err != nil {
			fmt.Println(err)
		}

		err = db.Queries.RemoveGamesFromGameSection(db.Ctx, c.Params("sectionid"))
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error removing games from section: %v", err))
		}

		uniqueGameIDs := []string{}
		for _, gameID := range payload.GameIDs {
			if slices.Index(uniqueGameIDs, gameID) == -1 {
				uniqueGameIDs = append(uniqueGameIDs, gameID)
			}
		}

		for i, gameID := range uniqueGameIDs {
			_, err = db.Queries.AddGameToGameSection(db.Ctx, dbpkg.AddGameToGameSectionParams{
				GameSectionID:  c.Params("sectionid"),
				GameID:         gameID,
				OrderInSection: int64(i),
			})
			if err != nil {
				return c.Status(500).SendString(fmt.Sprintf("Error adding game to section: %v", err))
			}
		}

		updatedGameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("pages/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	// Delete game from section
	router.Delete("/:sectionid/:gameid", func(c *fiber.Ctx) error {

		_, err := db.Queries.RemoveGameFromGameSection(db.Ctx, dbpkg.RemoveGameFromGameSectionParams{
			GameSectionID: c.Params("sectionid"),
			GameID:        c.Params("gameid"),
		})
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error deleting game from section: %v", err))
		}

		updatedGameSections, err := db.Queries.GetGameSectionsWithGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game sections with games: %v", err))
		}

		return c.Render("pages/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Get("/dialog/add", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/games/dialogsection", fiber.Map{
			"Edit": false,
		})
	})

	router.Get("/dialog/edit/:id", func(c *fiber.Ctx) error {

		sectionToEdit, err := db.Queries.GetGameSection(db.Ctx, c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game section: %v", err))
		}

		return c.Render("pages/admin/games/dialogsection", fiber.Map{
			"Edit":    true,
			"Section": sectionToEdit,
		})
	})
}
