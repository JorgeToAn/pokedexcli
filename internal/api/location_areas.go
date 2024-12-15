package api

import (
	"encoding/json"
	"io"
	"net/http"
)

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

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
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
