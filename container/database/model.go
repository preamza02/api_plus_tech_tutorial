package database

import (
	"github.com/uptrace/bun"
)

type PokemonModel struct {
	bun.BaseModel `bun:"table:pokemon,alias:p"`
	ID            int64              `bun:"id,pk,autoincrement"`
	Name          string             `bun:"name,notnull"`
	Description   *string            `bun:"description"`
	Category      *string            `bun:"category"`
	Types         []PokemonTypeModel `bun:"m2m:pokemon_types,join:Pokemon=PokemonType"`
	Abilities     []string           `bun:"type:array"`
}

type PokemonTypeModel struct {
	bun.BaseModel `bun:"table:pokemon_type,alias:pt"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
}

type PokemonTypesModel struct {
	bun.BaseModel `bun:"table:pokemon_types,alias:ppt"`
	PokemonID     int64             `bun:"pokemon_id,pk"`
	PokemonTypeID int64             `bun:"type_id,pk"`
	Pokemon       *PokemonModel     `bun:"rel:belongs-to,join:pokemon_id=id"`
	PokemonType   *PokemonTypeModel `bun:"rel:belongs-to,join:type_id=id"`
}
