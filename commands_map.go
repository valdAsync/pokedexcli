package main

import "fmt"

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if c.Next != nil && len(*c.Next) > 0 {
		url = *c.Next
	}

	locationAreas, err := getLocationAreas(url)
	if err != nil {
		return err
	}

	c.Next = &locationAreas.Next
	c.Previous = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	if c.Previous == nil || len(*c.Previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := getLocationAreas(*c.Previous)
	if err != nil {
		return err
	}

	c.Next = &locationAreas.Next
	c.Previous = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

