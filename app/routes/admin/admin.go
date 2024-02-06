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
		db.Preload("Games").Order("order_on_page").Find(&gameSections)
		var allGames []dbpkg.Game
		db.Order("title COLLATE NOCASE").Find(&allGames)
		return c.Render("routes/admin/index", fiber.Map{
			"Sections": gameSections,
			"Games":    allGames,
		}, "layouts/admin", "layouts/base")
	})

	router.Get("/games/sections", func(c *fiber.Ctx) error {
		var gameSections []dbpkg.GameSection
		db.Preload("Games").Order("order_on_page").Find(&gameSections)
		return c.Render("routes/admin/games/index", fiber.Map{
			"Sections": gameSections,
		})
	})

	router.Post("/games/sections", func(c *fiber.Ctx) error {

		var sectionsCount int64
		db.Model(&dbpkg.GameSection{}).Count(&sectionsCount)

		newSection := dbpkg.GameSection{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			OrderOnPage: int(sectionsCount + 1),
		}

		result := db.Create(&newSection)

		fmt.Println(result)

		var updatedGameSections []dbpkg.GameSection
		db.Preload("Games").Order("order_on_page").Find(&updatedGameSections)

		return c.Render("routes/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Put("/games/sections", func(c *fiber.Ctx) error {

		payload := struct {
			SectionIDs []string `form:"section"`
		}{}

		err := c.BodyParser(&payload)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(payload)

		for i, sectionid := range payload.SectionIDs {
			var sectionToUpdate dbpkg.GameSection
			db.First(&sectionToUpdate, "id", sectionid)
			sectionToUpdate.OrderOnPage = i
			db.Save(&sectionToUpdate)
		}

		var updatedGameSections []dbpkg.GameSection
		db.Preload("Games").Order("order_on_page").Find(&updatedGameSections)

		return c.Render("routes/admin/games/sectionslist", fiber.Map{
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
		db.Preload("Games").Order("order_on_page").Find(&updatedGameSections)

		return c.Render("routes/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Delete("/games/sections/:id", func(c *fiber.Ctx) error {
		var sectionToDelete dbpkg.GameSection
		db.First(&sectionToDelete, "id", c.Params("id"))
		db.Delete(&sectionToDelete)

		var updatedGameSections []dbpkg.GameSection
		db.Preload("Games").Order("order_on_page").Find(&updatedGameSections)

		return c.Render("routes/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Put("/games/sections/:sectionid/games", func(c *fiber.Ctx) error {

		payload := struct {
			GameIDs []string `form:"game"`
		}{}

		err := c.BodyParser(&payload)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(payload)

		var gameSectionGames []dbpkg.GameSectionGames
		db.Where("game_section_id = ?", c.Params("sectionid")).Delete(&gameSectionGames)

		for i, gameid := range payload.GameIDs {
			var gameSectionGames dbpkg.GameSectionGames

			gameSectionGames.GameSectionID = c.Params("sectionid")
			gameSectionGames.GameID = gameid
			gameSectionGames.OrderInSection = i

			db.Save(&gameSectionGames)
		}

		var updatedGameSections []dbpkg.GameSection
		db.Preload("Games").Order("order_on_page").Find(&updatedGameSections)

		return c.Render("routes/admin/games/sectionslist", fiber.Map{
			"Sections": updatedGameSections,
		})
	})

	router.Post("/games", func(c *fiber.Ctx) error {

		releaseYear, err := strconv.Atoi(c.FormValue("release_year"))
		if err != nil {
			return c.Status(400).SendString("Invalid release year, must be an integer")
		}

		rating, err := strconv.Atoi(c.FormValue("rating"))
		if err != nil {
			return c.Status(400).SendString("Invalid rating, must be an integer")
		}

		newGame := dbpkg.Game{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			Genre:       c.FormValue("genre"),
			ReleaseYear: releaseYear,
			ImageURL:    c.FormValue("image_url"),
			Rating:      rating,
		}

		db.Create(&newGame)

		var allGames []dbpkg.Game
		db.Order("title COLLATE NOCASE").Find(&allGames)

		return c.Render("routes/admin/games/gameslist", fiber.Map{
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

		rating, err := strconv.Atoi(c.FormValue("rating"))
		if err != nil {
			return c.Status(400).SendString("Invalid rating, must be an integer")
		}

		gameToUpdate.Title = c.FormValue("title")
		gameToUpdate.Description = c.FormValue("description")
		gameToUpdate.Genre = c.FormValue("genre")
		gameToUpdate.ReleaseYear = releaseYear
		gameToUpdate.ImageURL = c.FormValue("image_url")
		gameToUpdate.Rating = rating

		db.Save(&gameToUpdate)

		var allGames []dbpkg.Game
		db.Order("title COLLATE NOCASE").Find(&allGames)

		//remember to create an event or something to update the games in the other list

		return c.Render("routes/admin/games/gameslist", fiber.Map{
			"Games": allGames,
		})
	})

	router.Delete("/games/:id", func(c *fiber.Ctx) error {
		var gameToDelete dbpkg.Game
		db.First(&gameToDelete, "id", c.Params("id"))
		db.Delete(&gameToDelete)

		var allGames []dbpkg.Game
		db.Order("title COLLATE NOCASE").Find(&allGames)

		return c.Render("routes/admin/games/gameslist", fiber.Map{
			"Games": allGames,
		})
	})

	router.Get("/games/sections/dialog/add", func(c *fiber.Ctx) error {
		return c.Render("routes/admin/games/sectionsdialog", fiber.Map{
			"Edit": false,
		})
	})

	router.Get("/games/sections/dialog/edit/:id", func(c *fiber.Ctx) error {
		var sectionToEdit dbpkg.GameSection
		db.First(&sectionToEdit, "id", c.Params("id"))
		return c.Render("routes/admin/games/sectionsdialog", fiber.Map{
			"Edit":    true,
			"Section": sectionToEdit,
		})
	})

	router.Get("/games/dialog/add", func(c *fiber.Ctx) error {
		return c.Render("routes/admin/games/gamedialog", fiber.Map{
			"Edit": false,
		})
	})

	router.Get("/games/dialog/edit/:id", func(c *fiber.Ctx) error {
		var gameToEdit dbpkg.Game
		db.First(&gameToEdit, "id", c.Params("id"))
		return c.Render("routes/admin/games/gamedialog", fiber.Map{
			"Edit": true,
			"Game": gameToEdit,
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
