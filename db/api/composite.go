package db

import (
	db "gopplista/db/api/gen"
)

type SectionWithGames struct {
	db.GameSection
	Games []db.Game
}

func (q *db.Queries) GetSectionsWithGames() ([]SectionWithGames, error) {
	sections, err := db.Queries.GetGameSections(db.Ctx)
	if err != nil {
		return nil, err
	}

	sectionsWithGames := []SectionWithGames{}

	for _, section := range sections {
		games, err := db.Queries.GetGamesInGameSection(db.Ctx, section.ID)
		if err != nil {
			return nil, err
		}
		sectionsWithGames = append(sectionsWithGames, SectionWithGames{section, games})
	}

	return sectionsWithGames, nil
}
