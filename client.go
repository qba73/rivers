package rivers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	baseurl = "http://waterlevel.ie"
)

type errResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *client {
	return &client{
		BaseURL: baseurl,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) GetStations() (Stations, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/geojson/", c.BaseURL), nil)
	if err != nil {
		return Stations{}, err
	}

	var s Stations

	if err := c.sendRequest(req, &s); err != nil {
		return Stations{}, err
	}

	return s, nil
}

func (c *client) GetLatest() (StationsLatest, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/geojson/latest", c.BaseURL), nil)
	if err != nil {
		return StationsLatest{}, err
	}

	var s StationsLatest

	if err := c.sendRequest(req, &s); err != nil {
		return StationsLatest{}, err
	}

	return s, nil
}

func (c *client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return errors.New(errRes.Message)
		}
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
