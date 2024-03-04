-- GAME SECTIONS_LIST
-- name: unsafeCreateGameSection :one
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
   description = ?
WHERE
   id = ? RETURNING *;

-- name: unsafeUpdateGameSectionOrder :one
UPDATE game_sections
SET
   order_on_page = ?
WHERE
   id = ? RETURNING *;

-- name: unsafeDeleteGameSection :one
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
   title COLLATE NOCASE;

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
-- name: AddGameToGameSection :one
INSERT INTO
   game_section_games (game_id, game_section_id, order_in_section)
VALUES
   (?, ?, ?) RETURNING *;

-- name: GetGameSectionGames :many
SELECT
   *
FROM
   game_section_games
ORDER BY
   order_in_section;

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

-- name: RemoveGameFromGameSection :one
DELETE FROM game_section_games
WHERE
   game_section_id = ?
   AND game_id = ? RETURNING *;

-- name: RemoveGamesFromGameSection :exec
DELETE FROM game_section_games
WHERE
   game_section_id = ?;