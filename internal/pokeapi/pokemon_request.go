package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(pokemonName *string) (Pokemon, error) {
	if pokemonName == nil {
		return Pokemon{}, errors.New("put an id or name of an area")
	}
	endpoint := "/pokemon/" + *pokemonName
	fullUrl := BaseUrl + endpoint

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemon, nil
}
