package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetLocationAreaDetail(areaName string) (LocationAreaDetailResponse, error) {
	requestURL := baseURL + "/location-area/" + areaName

	value, exists := c.cache.Get(requestURL)
	if exists {
		result := LocationAreaDetailResponse{}
		err := json.Unmarshal(value, &result)
		if err != nil {
			return LocationAreaDetailResponse{}, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}

	result := LocationAreaDetailResponse{}
	if err = json.Unmarshal(data, &result); err != nil {
		return LocationAreaDetailResponse{}, err
	}

	c.cache.Add(requestURL, data)

	return result, nil
}
