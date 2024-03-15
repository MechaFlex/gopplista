package routes

import (
	"fmt"
	dbpkg "gopplista/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterGamesGamesRoutes(router fiber.Router, db dbpkg.Database) {

	router.Post("/", func(c *fiber.Ctx) error {

		rating, err := strconv.Atoi(c.FormValue("rating"))
		if err != nil {
			return c.Status(400).SendString("Invalid rating, must be an integer")
		}

		releaseYear, err := strconv.Atoi(c.FormValue("release_year"))
		if err != nil {
			return c.Status(400).SendString("Invalid release year, must be an integer")
		}

		db.Queries.CreateGame(db.Ctx, dbpkg.CreateGameParams{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Genre:       c.FormValue("genre"),
			ReleaseYear: int64(releaseYear),
			Rating:      int64(rating),
			ImageUrl:    c.FormValue("image_url"),
		})

		allGames, err := db.Queries.GetGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting games: %v", err))
		}

		return c.Render("pages/admin/games/gameslist", fiber.Map{
			"Games": allGames,
		})
	})

	router.Put("/:id", func(c *fiber.Ctx) error {

		rating, err := strconv.Atoi(c.FormValue("rating"))
		if err != nil {
			return c.Status(400).SendString("Invalid rating, must be an integer")
		}

		releaseYear, err := strconv.Atoi(c.FormValue("release_year"))
		if err != nil {
			return c.Status(400).SendString("Invalid release year, must be an integer")
		}

		db.Queries.UpdateGame(db.Ctx, dbpkg.UpdateGameParams{
			ID:          c.Params("id"),
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Genre:       c.FormValue("genre"),
			ReleaseYear: int64(releaseYear),
			Rating:      int64(rating),
			ImageUrl:    c.FormValue("image_url"),
		})

		//remember to create an event or something to update the games in the other list

		allGames, err := db.Queries.GetGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting games: %v", err))
		}

		c.Set("HX-Trigger", "game-updated")

		return c.Render("pages/admin/games/gameslist", fiber.Map{
			"Games": allGames,
		})
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {

		_, err := db.Queries.DeleteGame(db.Ctx, c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error deleting game: %v", err))
		}

		allGames, err := db.Queries.GetGames(db.Ctx)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting games: %v", err))
		}

		c.Set("HX-Trigger", "game-deleted")

		return c.Render("pages/admin/games/gameslist", fiber.Map{
			"Games": allGames,
		})
	})

	router.Get("/dialog/add", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/games/dialoggame", fiber.Map{
			"Edit": false,
		})
	})

	router.Get("/dialog/edit/:id", func(c *fiber.Ctx) error {

		gameToEdit, err := db.Queries.GetGame(db.Ctx, c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error getting game: %v", err))
		}

		return c.Render("pages/admin/games/dialoggame", fiber.Map{
			"Edit": true,
			"Game": gameToEdit,
		})
	})
}
