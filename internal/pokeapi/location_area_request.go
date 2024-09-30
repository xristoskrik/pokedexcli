package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullUrl := BaseUrl + endpoint
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("bad status code %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}
	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	c.cache.Add(fullUrl, data)
	return locationAreaResp, nil
}

func (c *Client) ListLocationAreasPokemon(area *string) (LocationAreaRespPokemon, error) {
	if area == nil {
		return LocationAreaRespPokemon{}, errors.New("put an id or name of an area")
	}
	endpoint := "/location-area/" + *area
	fullUrl := BaseUrl + endpoint

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationAreaRespPokemon{}, err
		}
		return LocationAreaRespPokemon{}, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaRespPokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRespPokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaRespPokemon{}, fmt.Errorf("bad status code %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRespPokemon{}, err
	}

	locationAreaRespPokemon := LocationAreaRespPokemon{}
	err = json.Unmarshal(data, &locationAreaRespPokemon)
	if err != nil {
		return LocationAreaRespPokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreaRespPokemon, nil
}
