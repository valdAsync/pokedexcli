package main

import (
	"fmt"
)

func commandPokedex(c *config, arg ...string) error {

	if len(c.pokedex) < 1 {
		fmt.Println("Your pokedex is empty! Go and catch them all!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, val := range c.pokedex {
		fmt.Printf(" - %s\n", val.Name)
	}

	return nil
}
