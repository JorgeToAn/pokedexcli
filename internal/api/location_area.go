package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c Client) GetLocationAreas(url *string) (LocationAreasResponse, error) {
	requestURL := baseURL + "/location-area"
	if url != nil {
		requestURL = *url
	}

	value, exists := c.cache.Get(requestURL)
	if exists {
		result := LocationAreasResponse{}
		if err := json.Unmarshal(value, &result); err != nil {
			return LocationAreasResponse{}, err
		}
		return result, nil
	}

	request, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	result := LocationAreasResponse{}
	if err = json.Unmarshal(data, &result); err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(requestURL, data)

	return result, nil
}
