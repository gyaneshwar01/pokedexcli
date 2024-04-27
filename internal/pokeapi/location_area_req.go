package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageurl *string) (LocationAreaResp, error) {
	var fullURL string

	if pageurl == nil {
		endpoint := "/location-area"
		fullURL = baseURL + endpoint
	} else {
		fullURL = *pageurl
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code %v", res.StatusCode)
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreasResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationAreasResp, nil
}
