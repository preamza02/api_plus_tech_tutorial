package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func InitSqlite() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:db/pokedex.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.RegisterModel((*PokemonTypesModel)(nil))
	return db
}
