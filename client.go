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

	if err := c.sendRequestJSON(req, &s); err != nil {
		return Stations{}, err
	}

	return s, nil
}

func (c *client) GetLatest() (StationsLatest, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geojson/latest", c.BaseURL), nil)
	if err != nil {
		return StationsLatest{}, err
	}

	var s StationsLatest

	if err := c.sendRequestJSON(req, &s); err != nil {
		return StationsLatest{}, err
	}

	return s, nil
}

func (c *client) GetDayLevel(stationID string) ([]SensorReading, error) {
	url := c.urlLevel(stationID, "day")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

func (c *client) GetWeekLevel(stationID string) ([]SensorReading, error) {
	url := c.urlLevel(stationID, "week")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

func (c *client) GetMonthLevel(stationID string) ([]SensorReading, error) {
	url := c.urlLevel(stationID, "month")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

func (c *client) GetDayTemperature(statioID string) ([]SensorReading, error) {
	url := c.urlTemperature(statioID, "day")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

func (c *client) GetWeekTemperature(stationID string) ([]SensorReading, error) {
	url := c.urlTemperature(stationID, "week")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

func (c *client) GetMonthTemperature(stationID string) ([]SensorReading, error) {
	url := c.urlTemperature(stationID, "month")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

func (c *client) urlLevel(stationID, period string) string {
	var url string
	switch period {
	case "day":
		url = fmt.Sprintf("%s/data/day/%s", c.BaseURL, fileNameLevel(stationID))
	case "week":
		url = fmt.Sprintf("%s/data/week/%s", c.BaseURL, fileNameLevel(stationID))
	case "month":
		url = fmt.Sprintf("%s/data/month/%s", c.BaseURL, fileNameLevel(stationID))
	}

	return url
}

func (c *client) urlTemperature(stationID, period string) string {
	var url string
	switch period {
	case "day":
		url = fmt.Sprintf("%s/data/day/%s", c.BaseURL, fileNameTemperature(stationID))
	case "week":
		url = fmt.Sprintf("%s/data/week/%s", c.BaseURL, fileNameTemperature(stationID))
	case "month":
		url = fmt.Sprintf("%s/data/month/%s", c.BaseURL, fileNameTemperature(stationID))
	}

	return url
}

func (c *client) urlVoltage(stationID, period string) string {
	var url string
	switch period {
	case "day":
		url = fmt.Sprintf("%s/data/day/%s", c.BaseURL, fileNameVoltage(stationID))
	case "week":
		url = fmt.Sprintf("%s/data/week/%s", c.BaseURL, fileNameVoltage(stationID))
	case "month":
		url = fmt.Sprintf("%s/data/day/%s", c.BaseURL, fileNameVoltage(stationID))
	}

	return url
}

func (c *client) sendRequestJSON(req *http.Request, v interface{}) error {
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

func (c *client) sendRequestCSV(req *http.Request) ([]SensorReading, error) {
	req.Header.Set("Content-Type", "text/csv")
	req.Header.Set("Accept", "text/csv")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	out, err := ReadCSV(res.Body)
	if err != nil {
		return nil, err
	}

	return out, nil

}

func fileNameLevel(stationID string) string {
	return fmt.Sprintf("%s_000%v.csv", stationID, levelSensor)
}

func fileNameTemperature(stationID string) string {
	return fmt.Sprintf("%s_000%v.csv", stationID, tempSensor)
}

func fileNameVoltage(stationID string) string {
	return fmt.Sprintf("%s_000%v.csv", stationID, voltageSensor)
}

func fileNameOD(stationID string) string {
	return fmt.Sprintf("%s_OD.csv", stationID)
}
