package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	// request
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	// response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	// for 400 (bad req), 500 (server err) responses
	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body) // read the data
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{} // where we store the data
	// unmarshalling the data: taking a pointer to a structure and filling in the data
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResp, nil
}
