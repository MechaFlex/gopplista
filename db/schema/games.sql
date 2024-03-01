DROP TABLE IF EXISTS games;

CREATE TABLE
   games (
      id TEXT NOT NULL PRIMARY KEY DEFAULT (HEX (RANDOMBLOB (8))),
      title TEXT NOT NULL,
      description TEXT NOT NULL,
      genre TEXT NOT NULL,
      release_year INT NOT NULL,
      rating INT NOT NULL,
      image_url TEXT NOT NULL
   );

DROP TABLE IF EXISTS game_sections;

CREATE TABLE
   game_sections (
      id TEXT NOT NULL PRIMARY KEY DEFAULT (HEX (RANDOMBLOB (8))),
      title TEXT NOT NULL,
      description TEXT NOT NULL,
      order_on_page INT NOT NULL,
      UNIQUE (order_on_page)
   );

DROP TABLE IF EXISTS game_section_games;

CREATE TABLE
   game_section_games (
      game_section_id TEXT NOT NULL,
      game_id TEXT NOT NULL,
      order_in_section INT NOT NULL,
      PRIMARY KEY (game_section_id, game_id),
      FOREIGN KEY (game_section_id) REFERENCES game_sections (id),
      FOREIGN KEY (game_id) REFERENCES games (id),
      UNIQUE (game_section_id, order_in_section)
   );