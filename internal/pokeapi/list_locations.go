package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreas, error) {
	url := baseUrl + "location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}
	var locationAreas LocationAreas

	val, exists := c.cache.Get(url)
	if exists {
		if err := json.Unmarshal(val, &locationAreas); err == nil {
			return locationAreas, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, body)
	return locationAreas, nil
}
