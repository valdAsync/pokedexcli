package main

import "fmt"

func commandMap(c *config) error {
	locationAreas, err := c.pokeapiClient.ListLocationAreas(c.nextLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locationAreas.Next
	c.prevLocationURL = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	if c.nextLocationURL == nil || len(*c.prevLocationURL) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := c.pokeapiClient.ListLocationAreas(c.nextLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locationAreas.Next
	c.prevLocationURL = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

