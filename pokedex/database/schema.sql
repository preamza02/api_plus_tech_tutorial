Create Table pokemon (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL
);

Create Table pokemon_type{
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
}
 
Create Table pokemon_types{
   pokemon_id INTEGER FOREIGN KEY REFERENCES pokemon(id),
   type_id INTEGER FOREIGN KEY REFERENCES pokemon_type(id)
}

Insert into pokemon_type (name) values 
("NORMAL"),("FIRE"),("WATER"),("ICE"),("FIGHTING")