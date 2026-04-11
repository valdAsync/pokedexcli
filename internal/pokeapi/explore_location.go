package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocationArea(areaName string) (ExploredLocationArea, error) {
	url := baseUrl + "location-area/" + areaName

	var exploredLocationArea ExploredLocationArea

	val, exists := c.cache.Get(url)
	if exists {
		if err := json.Unmarshal(val, &exploredLocationArea); err == nil {
			return exploredLocationArea, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return ExploredLocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploredLocationArea{}, err
	}

	if err := json.Unmarshal(body, &exploredLocationArea); err != nil {
		return ExploredLocationArea{}, err
	}

	c.cache.Add(url, body)
	return exploredLocationArea, nil
}
