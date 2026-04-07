package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreas, error) {
	url := baseUrl + "location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}

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
