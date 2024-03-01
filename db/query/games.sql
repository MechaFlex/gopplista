-- GAME SECTIONS
-- name: CreateGameSection :one
INSERT INTO
   game_sections (title, description, order_on_page)
VALUES
   (?, ?, ?) RETURNING *;

-- name: GetGameSections :many
SELECT
   *
FROM
   game_sections
ORDER BY
   order_on_page;

-- name: GetGameSection :one
SELECT
   *
FROM
   game_sections
WHERE
   id = ?;

-- name: UpdateGameSection :one
UPDATE game_sections
SET
   title = ?,
   description = ?,
   order_on_page = ?
WHERE
   id = ? RETURNING *;

-- name: DeleteGameSection :one
DELETE FROM game_sections
WHERE
   id = ? RETURNING *;

-- name: CountGameSections :one
SELECT
   COUNT(*)
FROM
   game_sections;

-- GAMES
-- name: CreateGame :one
INSERT INTO
   games (
      title,
      description,
      genre,
      release_year,
      rating,
      image_url
   )
VALUES
   (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: GetGames :many
SELECT
   *
FROM
   games
ORDER BY
   title;

-- name: GetGame :one
SELECT
   *
FROM
   games
WHERE
   id = ?;

-- name: UpdateGame :one
UPDATE games
SET
   title = ?,
   description = ?,
   genre = ?,
   release_year = ?,
   rating = ?,
   image_url = ?
WHERE
   id = ? RETURNING *;

-- name: DeleteGame :one
DELETE FROM games
WHERE
   id = ? RETURNING *;

-- GAME SECTION'S GAMES
-- name: AddGameToSection :one
INSERT INTO
   game_section_games (game_id, game_section_id, order_in_section)
VALUES
   (?, ?, ?) RETURNING *;

-- name: GetGamesInGameSection :many
SELECT
   games.*
FROM
   game_section_games
   JOIN games ON game_section_games.game_id = games.id
WHERE
   game_section_id = ?
ORDER BY
   order_in_section;

-- name: GetAllGameSectionsWithGames :many
SELECT
   game_sections.*, sqlc.embed(games)
FROM
   game_sections
   LEFT JOIN game_section_games ON game_sections.id = game_section_games.game_section_id
   LEFT JOIN games ON game_section_games.game_id = games.id
ORDER BY
   game_sections.order_on_page,
   game_section_games.order_in_section;

-- name: updateGameInSectionOrder :one
UPDATE game_section_games
SET
   order_in_section = ?
WHERE
   game_id = ?
   AND game_section_id = ? RETURNING *;

-- name: RemoveGameFromSection :one
DELETE FROM game_section_games
WHERE
   game_id = ?
   AND game_section_id = ? RETURNING *;