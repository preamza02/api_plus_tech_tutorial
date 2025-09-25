package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
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

func InitPostgres() *bun.DB {
	dsn := "postgres://pokedex:password@postgres:5432/pokedex_db?sslmode=disable"
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open Postgres connection:", err)
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	db.RegisterModel((*PokemonTypesModel)(nil))
	if err := sqldb.Ping(); err != nil {
		log.Fatal("Failed to connect to Postgres database:", err)
	}
	return db
}
