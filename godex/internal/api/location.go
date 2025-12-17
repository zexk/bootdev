package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationRes struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocation(location string) (LocationRes, error) {
	url := baseUrl + "location-area/" + location

	if data, exists := c.cache.Get(url); exists {
		locRes := LocationRes{}
		err := json.Unmarshal(data, &locRes)
		if err != nil {
			return LocationRes{}, err
		}
		return locRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationRes{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return LocationRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationRes{}, err
	}

	locRes := LocationRes{}
	err = json.Unmarshal(data, &locRes)
	if err != nil {
		return LocationRes{}, err
	}

	c.cache.Set(url, data)
	return locRes, nil
}
