package admin

import (
	"fmt"
	dbpkg "gopplista/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(router fiber.Router, db *gorm.DB) {

	// router.Use(func(c *fiber.Ctx) error {
	// 	if c.Cookies("admin") != "true" {
	// 		return c.Redirect("/login")
	// 	}
	// 	return c.Next()
	// })

	router.Get("/", func(c *fiber.Ctx) error {
		var gameSections []dbpkg.GameSection
		db.Preload("Games").Find(&gameSections)
		var allGames []dbpkg.Game
		db.Order("title DESC").Find(&allGames)
		return c.Render("routes/admin/index", fiber.Map{
			"Sections": gameSections,
			"Games":    allGames,
		}, "layouts/admin", "layouts/base")
	})

	router.Get("/games/sections", func(c *fiber.Ctx) error {
		var gameSections []dbpkg.GameSection
		db.Preload("Games").Find(&gameSections)
		return c.Render("routes/admin/games/index", fiber.Map{
			"Sections": gameSections,
		})
	})

	router.Post("/games/sections", func(c *fiber.Ctx) error {

		newSection := dbpkg.GameSection{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			OrderOnPage: 0,
		}

		result := db.Create(&newSection)

		fmt.Println(result)

		var updatedGameSections []dbpkg.GameSection
		db.Preload("Games").Find(&updatedGameSections)

		fmt.Println(updatedGameSections)

		return c.Render("routes/admin/games/sections", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Put("/games/sections/:id", func(c *fiber.Ctx) error {
		var sectionToUpdate dbpkg.GameSection
		db.First(&sectionToUpdate, "id", c.Params("id"))

		sectionToUpdate.Title = c.FormValue("title")
		sectionToUpdate.Description = c.FormValue("description")

		db.Save(&sectionToUpdate)

		var updatedGameSections []dbpkg.GameSection
		db.Preload("Games").Find(&updatedGameSections)

		return c.Render("routes/admin/games/sections", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Delete("/games/sections/:id", func(c *fiber.Ctx) error {
		var sectionToDelete dbpkg.GameSection
		db.First(&sectionToDelete, "id", c.Params("id"))
		db.Delete(&sectionToDelete)

		var updatedGameSections []dbpkg.GameSection
		db.Preload("Games").Find(&updatedGameSections)

		return c.Render("routes/admin/games/sections", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Post("/games", func(c *fiber.Ctx) error {

		releaseYear, err := strconv.Atoi(c.FormValue("release_year"))
		if err != nil {
			return c.Status(400).SendString("Invalid release year, must be an integer")
		}

		newGame := dbpkg.Game{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Genre:       c.FormValue("genre"),
			ReleaseYear: releaseYear,
			ImageURL:    c.FormValue("image_url"),
			Score:       0,
		}

		db.Create(&newGame)

		var allGames []dbpkg.Game
		db.Order("title DESC").Find(&allGames)

		return c.Render("routes/admin/games/index", fiber.Map{
			"Games": allGames,
		})
	})

	router.Put("/games/:id", func(c *fiber.Ctx) error {
		var gameToUpdate dbpkg.Game
		db.First(&gameToUpdate, "id", c.Params("id"))

		releaseYear, err := strconv.Atoi(c.FormValue("release_year"))
		if err != nil {
			return c.Status(400).SendString("Invalid release year, must be an integer")
		}

		gameToUpdate.Title = c.FormValue("title")
		gameToUpdate.Description = c.FormValue("description")
		gameToUpdate.Genre = c.FormValue("genre")
		gameToUpdate.ReleaseYear = releaseYear
		gameToUpdate.ImageURL = c.FormValue("image_url")

		db.Save(&gameToUpdate)

		var allGames []dbpkg.Game
		db.Order("title DESC").Find(&allGames)

		//remember to create an event or something to update the games in the other list

		return c.Render("routes/admin/games/index", fiber.Map{
			"Games": allGames,
		})
	})

	router.Delete("/games/:id", func(c *fiber.Ctx) error {
		var gameToDelete dbpkg.Game
		db.First(&gameToDelete, "id", c.Params("id"))
		db.Delete(&gameToDelete)

		var allGames []dbpkg.Game
		db.Order("title DESC").Find(&allGames)

		return c.Render("routes/admin/games/index", fiber.Map{
			"Games": allGames,
		})
	})

	// router.Post("/games/sections/:sectionid/games", func(c *fiber.Ctx) error {

	// 	fmt.Println(c.Params("sectionid"))

	// 	releaseYear, err := strconv.Atoi(c.FormValue("release_year"))
	// 	if err != nil {
	// 		return c.Status(400).SendString("Invalid release year, must be an ineteger")
	// 	}

	// 	newGame := dbpkg.Game{
	// 		Title:       c.FormValue("title"),
	// 		Description: c.FormValue("description"),
	// 		Genre:       c.FormValue("genre"),
	// 		ReleaseYear: releaseYear,
	// 		ImageURL:    c.FormValue("image_url"),
	// 		Score:       0,
	// 	}
	// 	db.Save(&newGame)

	// 	var sectionToUpdate dbpkg.GameSection
	// 	db.First(&sectionToUpdate, "id", c.Params("sectionid"))
	// 	db.Preload("Games").Find(&sectionToUpdate)

	// 	sectionToUpdate.Games = append(sectionToUpdate.Games, newGame)
	// 	db.Save(&sectionToUpdate)

	// 	var updatedGameSections []dbpkg.GameSection
	// 	db.Preload("Games").Find(&updatedGameSections)

	// 	return c.Render("routes/admin/games/sections", fiber.Map{
	// 		"Sections": updatedGameSections,
	// 	})
	// })
}
