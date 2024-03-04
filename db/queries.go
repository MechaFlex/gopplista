package db

import "context"

type CreateGameSectionParams struct {
	Title       string
	Description string
}

func (q *Queries) CreateGameSection(ctx context.Context, arg CreateGameSectionParams) (GameSection, error) {

	gameSectionCount, err := q.CountGameSections(ctx)
	if err != nil {
		return GameSection{}, err
	}

	result, err := q.unsafeCreateGameSection(ctx, unsafeCreateGameSectionParams{
		Title:       arg.Title,
		Description: arg.Description,
		OrderOnPage: gameSectionCount + 1,
	})
	if err != nil {
		return GameSection{}, err
	}

	return result, nil
}

type UpdateGameSectionOrderParams struct {
	ID          string
	OrderOnPage int64
}

func (q *Queries) UpdateGameSectionOrder(ctx context.Context, arg UpdateGameSectionOrderParams) (GameSection, error) {

	allGameSections, err := q.GetGameSections(ctx)
	if err != nil {
		return GameSection{}, err
	}

	for _, section := range allGameSections {
		if section.OrderOnPage >= arg.OrderOnPage {
			_, err = q.unsafeUpdateGameSectionOrder(ctx, unsafeUpdateGameSectionOrderParams{
				OrderOnPage: section.OrderOnPage + 1,
				ID:          section.ID,
			})
			if err != nil {
				return GameSection{}, err
			}
		}
	}

	result, err := q.unsafeUpdateGameSectionOrder(ctx, unsafeUpdateGameSectionOrderParams{
		OrderOnPage: arg.OrderOnPage,
		ID:          arg.ID,
	})
	if err != nil {
		return GameSection{}, err
	}

	err = settleGameSectionOrdering(q, ctx)
	if err != nil {
		return GameSection{}, err
	}

	return result, nil
}

func (q *Queries) DeleteGameSection(ctx context.Context, arg GameSection) (GameSection, error) {
	result, err := q.unsafeDeleteGameSection(ctx, arg.ID)
	if err != nil {
		return GameSection{}, err
	}

	err = settleGameSectionOrdering(q, ctx)
	if err != nil {
		return GameSection{}, err
	}

	return result, nil
}

func settleGameSectionOrdering(q *Queries, ctx context.Context) error {
	sections, err := q.GetGameSections(ctx)
	if err != nil {
		return err
	}

	for i, section := range sections {
		if section.OrderOnPage != int64(i) {
			_, err = q.unsafeUpdateGameSectionOrder(ctx, unsafeUpdateGameSectionOrderParams{
				OrderOnPage: int64(i),
				ID:          section.ID,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type SectionWithGames struct {
	GameSection
	Games []Game
}

func (q *Queries) GetGameSectionsWithGames(ctx context.Context) ([]SectionWithGames, error) {
	sections, err := q.GetGameSections(ctx)
	if err != nil {
		return nil, err
	}

	sectionsWithGames := []SectionWithGames{}

	for _, section := range sections {
		games, err := q.GetGamesInGameSection(ctx, section.ID)
		if err != nil {
			return nil, err
		}
		sectionsWithGames = append(sectionsWithGames, SectionWithGames{section, games})
	}

	return sectionsWithGames, nil
}
