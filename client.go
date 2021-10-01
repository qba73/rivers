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

// pattern ! make note in my notebook
var validPeriod = map[string]bool{
	"day":   true,
	"week":  true,
	"month": true,
}

type errResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient knows how to construct a new default rivers client.
// The client will be used to retrieve information about
// various measures recorded by sensors.
func NewClient() *client {
	return &client{
		BaseURL: baseurl,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetStations knows how to return information about all gauges
// (measurement stations) installed in rivers in Ireland.
func (c *client) GetStations() (Stations, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geojson/", c.BaseURL), nil)
	if err != nil {
		return Stations{}, err
	}

	var s Stations

	if err := c.sendRequestJSON(req, &s); err != nil {
		return Stations{}, err
	}

	return s, nil
}

// GetLatest knows how to return latest readings from water level,
// water temperature and voltage level sensors installed in all
// stations in rivers in Ireland.
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

// GetDayLevel knows how to return water level readings recorded for
// last 24hr period for the given stationID number.
func (c *client) GetDayLevel(stationID string) ([]SensorReading, error) {
	url, err := c.urlLevel(stationID, "day")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

// GetWeekLevel knows how to return water level readings recorded for
// last week period for the given stationID number.
func (c *client) GetWeekLevel(stationID string) ([]SensorReading, error) {
	url, err := c.urlLevel(stationID, "week")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

// GetMonthLevel knows how to return water level readings recorded for
// last 4 weeks period for the given stationID number.
func (c *client) GetMonthLevel(stationID string) ([]SensorReading, error) {
	url, err := c.urlLevel(stationID, "month")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

// GetDayTemperature knows how to return water temperature
// recorded for last 24hr period for the given stationID number.
func (c *client) GetDayTemperature(statioID string) ([]SensorReading, error) {
	url, err := c.urlTemperature(statioID, "day")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

// GetWeekTemperature knows how to return water temperature
// recorded for last week period for the given stationID number.
func (c *client) GetWeekTemperature(stationID string) ([]SensorReading, error) {
	url, err := c.urlTemperature(stationID, "week")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

// GetMonthTemperature knows how to return water temperature
// recorded for last 4 weeks period for the given stationID number.
func (c *client) GetMonthTemperature(stationID string) ([]SensorReading, error) {
	url, err := c.urlTemperature(stationID, "month")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

func (c *client) GetDayVoltage(stationID string) ([]SensorReading, error) {
	url, err := c.urlVoltage(stationID, "day")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequestCSV(req)
}

// urlLevel takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *client) urlLevel(stationID, period string) (string, error) {
	if !validPeriod[period] {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s", c.BaseURL, period, fileNameLevel(stationID)), nil
}

// urlTemperature takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *client) urlTemperature(stationID, period string) (string, error) {
	if !validPeriod[period] {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s", c.BaseURL, period, fileNameTemperature(stationID)), nil
}

// urlVoltage takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *client) urlVoltage(stationID, period string) (string, error) {
	if !validPeriod[period] {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s", c.BaseURL, period, fileNameVoltage(stationID)), nil
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

/*
func fileNameOD(stationID string) string {
	return fmt.Sprintf("%s_OD.csv", stationID)
}
*/
