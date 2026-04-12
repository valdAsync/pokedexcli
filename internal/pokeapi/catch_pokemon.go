package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (PokemonDetails, error) {
	url := baseUrl + "pokemon/" + pokemonName

	var pokemonDetails PokemonDetails

	res, err := http.Get(url)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	if err := json.Unmarshal(body, &pokemonDetails); err != nil {
		return PokemonDetails{}, err
	}

	return pokemonDetails, nil
}
