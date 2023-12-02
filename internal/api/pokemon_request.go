package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseUrl + endpoint

	// check the cache
	data, ok := c.cache.Get(fullUrl)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		pokemon := Pokemon{} // where we store the data
		// unmarshalling the data: taking a pointer to a structure and filling in the data
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil // if there is a cache hit, return successfully early
	}
	fmt.Println("cache miss")

	// request
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	// for 400 (bad req), 500 (server err) responses
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body) // read the data
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{} // where we store the data
	// unmarshalling the data: taking a pointer to a structure and filling in the data
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemon, nil
}
