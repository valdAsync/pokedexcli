package main

import (
	"fmt"
)

func commandInspect(c *config, arg ...string) error {
	if len(arg) < 1 {
		return fmt.Errorf("Inspect command requires a pokemon name")
	}

	val, ok := c.pokedex[arg[0]]

	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Weight: %v\n", val.Weight)
	fmt.Printf("Height: %v\n", val.Height)

	fmt.Println("Stats:")
	for _, statsVals := range val.Stats {
		fmt.Printf("  -%s: %v\n", statsVals.Stat.Name, statsVals.BaseStat)
	}

	fmt.Println("Types:")
	for _, typesVals := range val.Types {
		fmt.Printf("  - %s\n", typesVals.Type.Name)
	}

	return nil
}
