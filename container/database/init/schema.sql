CREATE TABLE
   pokemon (
      id INTEGER PRIMARY KEY,
      name TEXT UNIQUE NOT NULL,
      description TEXT,
      category TEXT,
      abilities TEXT
   );

CREATE TABLE
   pokemon_type (id INTEGER PRIMARY KEY, name TEXT NOT NULL);

CREATE TABLE
   pokemon_types (
      pokemon_id INTEGER,
      type_id INTEGER,
      FOREIGN KEY (pokemon_id) REFERENCES pokemon (id),
      FOREIGN KEY (type_id) REFERENCES pokemon_type (id),
      PRIMARY KEY (pokemon_id, type_id)
   );


INSERT INTO
   pokemon_type (id,name)
values
   (1,'NORMAL'),
   (2,'FIRE'),
   (3,'WATER'),
   (4,'ICE'),
   (5,'FIGHTING');

INSERT INTO
   pokemon (id, name, description, category, abilities)
values
   (
      1,
      'Bulbasaur',
      'A strange seed was planted on its back at birth',
      'something',
      '["Overgrow", "Chlorophyll"]'
   ),
   (
      2,
      'Charmander',
      'Obviously prefers hot places. When it rains, steam is said to spout from the tip of its tail.',
      'some some',
      '["Blaze"]'
   ),
   (
      3,
      'Squirtle',
      'After birth, its back swells and hardens into a shell. Powerfully sprays foam from its mouth.',
      'thing thing',
      '["Torrent"]'
   ),
   (
      4,
      'Pikachu',
      'When several of these Pokémon gather, their electricity could build and cause lightning storms.',
      'nice',
      '["Static", "Lightning Rod"]'
   );

INSERT INTO
   pokemon_types (pokemon_id, type_id)
values
   (1, 1),
   (1, 4),
   (2, 2),
   (3, 3),
   (4, 1);