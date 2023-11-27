package api

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"` // use pointer when string || nil
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
