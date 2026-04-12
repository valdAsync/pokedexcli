package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, arg ...string) error {
	if len(arg) < 1 {
		return fmt.Errorf("Catch command requires a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg[0])
	pokemonResponse, err := c.pokeapiClient.CatchPokemon(arg[0])
	if err != nil {
		return err
	}

	if caught := attemptCatch(pokemonResponse.BaseExperience); caught == true {
		c.pokedex[pokemonResponse.Name] = pokemonResponse
		fmt.Printf("%s was caught!\n", pokemonResponse.Name)
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonResponse.Name)

	return nil
}

func attemptCatch(pokemonExperience int) bool {
	myPower := rand.IntN(500)

	if myPower > pokemonExperience {
		return true
	}

	return false
}
