INSERT INTO
   games (
      title,
      description,
      genre,
      releaseYear,
      rating,
      imageURL
   )
VALUES
   (
      'The Legend of Zelda: Breath of the Wild',
      'The Legend of Zelda: Breath of the Wild is an action-adventure game developed and published by Nintendo, released for the Nintendo Switch and Wii U consoles on March 3, 2017. The game is a part of The Legend of Zelda series, and follows amnesiac protagonist Link, who awakens from a hundred-year slumber to a mysterious voice that guides him to defeat Calamity Ganon before he can destroy the kingdom of Hyrule.',
      'Action-adventure',
      2017,
      4,
      'https://images.igdb.com/igdb/image/upload/t_cover_big/co3p2d.png'
   );

INSERT INTO
   gameSections (title, description, orderOnPage)
VALUES
   (
      'Featured Games',
      'The best games we have to offer',
      0
   );

INSERT INTO
   gameSectionGames
VALUES
   (
      (
         SELECT
            id
         FROM
            gameSections
         WHERE
            title = 'Featured Games'
      ),
      (
         SELECT
            id
         FROM
            games
         WHERE
            title = 'The Legend of Zelda: Breath of the Wild'
      ),
      0
   );