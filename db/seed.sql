INSERT INTO
	games (
		id,
		title,
		description,
		genre,
		release_year,
		rating,
		image_url
	)
VALUES
	(
		'ID-botw',
		'The Legend of Zelda: Breath of the Wild',
		'The Legend of Zelda: Breath of the Wild is an action-adventure game developed and published by Nintendo, released for the Nintendo Switch and Wii U consoles on March 3, 2017. The game is a part of The Legend of Zelda series, and follows amnesiac protagonist Link, who awakens from a hundred-year slumber to a mysterious voice that guides him to defeat Calamity Ganon before he can destroy the kingdom of Hyrule.',
		'Action-adventure',
		2017,
		5,
		'https://images.igdb.com/igdb/image/upload/t_cover_big/co3p2d.png'
	),
	(
		'ID-animalcrossing',
		'Animal Crossing: New Horizons',
		'Animal Crossing: New Horizons is a 2020 life simulation video game developed and published by Nintendo for the Nintendo Switch; it is the fifth main series title in the Animal Crossing series. In New Horizons, the player assumes the role of a customizable character who moves to a deserted island after purchasing a package from Tom Nook, a raccoon character who is a staple of the series. Taking place in real-time, the player can explore the island in a nonlinear fashion, gathering and crafting items, catching insects and fish, developing the island into a community of anthropomorphic animals.',
		'Life simulation',
		2020,
		2,
		'https://images.igdb.com/igdb/image/upload/t_cover_big/co3wls.png'
	),
	(
		'ID-odyssey',
		'Super Mario Odyssey',
		'Super Mario Odyssey is a platform game developed and published by Nintendo for the Nintendo Switch on October 27, 2017. An entry in the Super Mario series, it follows Mario and Cappy, a sentient hat that allows Mario to control other characters and objects, as they journey across various worlds to save Princess Peach from his nemesis Bowser, who plans to forcibly marry her. In contrast to the linear gameplay of prior entries, the game returns to the primarily open-ended, exploration-based gameplay featured in Super Mario 64 and Super Mario Sunshine.',
		'3D platformer',
		2017,
		4,
		'https://images.igdb.com/igdb/image/upload/t_cover_big/co1mxf.png'
	),
	(
		'ID-celeste',
		'Celeste',
		'Celeste is a 2018 platform game developed and published by indie studio Maddy Makes Games. The player controls Madeline, a young woman with anxiety and depression who aims to climb Celeste Mountain. During her climb, she encounters several characters, including a personification of her self-doubt known as Badeline, who attempts to stop her from climbing the mountain.',
		'2D platformer',
		2018,
		4,
		'https://images.igdb.com/igdb/image/upload/t_cover_big/co3byy.png'
	);

INSERT INTO
	game_sections (id, title, description, order_on_page)
VALUES
	(
		'ID-nintendogames',
		'Nintendo Games',
		'Games developed and published by Nintendo',
		0
	),
	(
		'ID-platformers',
		'Platformers',
		'Games that are platformers',
		1
	);

INSERT INTO
	game_section_games
VALUES
	('ID-nintendogames', 'ID-botw', 0),
	('ID-nintendogames', 'ID-animalcrossing', 1),
	('ID-nintendogames', 'ID-odyssey', 2),
	('ID-platformers', 'ID-odyssey', 0),
	('ID-platformers', 'ID-celeste', 1);