package main

import "fmt"

func commandExplore(c *config, arg ...string) error {
	if len(arg) < 1 {
		return fmt.Errorf("Explore command requires an area name argument")
	}

	fmt.Printf("Exploring %s...\n", arg[0])
	exploredLocationArea, err := c.pokeapiClient.ExploreLocationArea(arg[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemons := range exploredLocationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemons.Pokemon.Name)
	}

	return nil
}
