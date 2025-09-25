build:
	go build -o bin/pokedex 

data/pokedex.db:
	mkdir -p data
	sqlite3 data/pokedex.db  < database/schema.sql