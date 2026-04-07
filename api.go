package main

import (
	"encoding/json"
	"net/http"
)

type LocationAreas struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous *string               `json:"previous"`
	Results  []LocationAreaResults `json:"results"`
}

type LocationAreaResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getLocationAreas(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	var locationAreas LocationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}

	return locationAreas, nil
}
