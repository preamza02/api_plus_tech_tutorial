package model

import (
	"fmt"

	"github.com/preamza02/pokedex/database"
)

func MapDatabasePokemonModelToPokemon(pokemon *database.PokemonModel) *Pokemon {
	answer := &Pokemon{
		ID:          fmt.Sprint(pokemon.ID),
		Name:        pokemon.Name,
		Description: pokemon.Description,
		Category:    pokemon.Category,
		Type:        make([]PokemonType, len(pokemon.Types)),
		Abilities:   pokemon.Abilities,
	}
	for i, t := range pokemon.Types {
		answer.Type[i] = PokemonType(t.Name)
	}
	return answer
}
