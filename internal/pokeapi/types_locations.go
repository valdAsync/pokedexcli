package pokeapi

type LocationAreas struct {
	Count    int                   `json:"count"`
	Next     *string               `json:"next"`
	Previous *string               `json:"previous"`
	Results  []LocationAreaResults `json:"results"`
}

type LocationAreaResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
