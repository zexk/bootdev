package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaRes struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) PrintLocationArea(pageUrl *string) (LocationAreaRes, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if data, exists := c.cache.Get(url); exists {
		locAreaRes := LocationAreaRes{}
		err := json.Unmarshal(data, &locAreaRes)
		if err != nil {
			return LocationAreaRes{}, err
		}
		return locAreaRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaRes{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	locAreaRes := LocationAreaRes{}
	err = json.Unmarshal(data, &locAreaRes)
	if err != nil {
		return LocationAreaRes{}, err
	}

	c.cache.Set(url, data)
	return locAreaRes, nil
}
