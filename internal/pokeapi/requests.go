package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(url)

	if ok {
		fmt.Println("cache hit!")

		locationAreasResponse := LocationAreasResponse{}

		err := json.Unmarshal(data, &locationAreasResponse)

		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResponse, nil
	}

	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}

	err = json.Unmarshal(data, &locationAreasResponse)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Set(url, data)

	return locationAreasResponse, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationAreaName

	// check the cache
	data, ok := c.cache.Get(url)

	if ok {
		fmt.Println("cache hit!")

		locationArea := LocationArea{}

		err := json.Unmarshal(data, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}

	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Set(url, data)

	return locationArea, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// check the cache
	data, ok := c.cache.Get(url)

	if ok {
		// fmt.Println("cache hit!")

		pokemon := Pokemon{}

		err := json.Unmarshal(data, &pokemon)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	// fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Set(url, data)

	return pokemon, nil
}
