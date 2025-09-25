package database

import (
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func InitSqlite() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:data/pokedex.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.RegisterModel((*PokemonTypesModel)(nil))
	if err := sqldb.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
