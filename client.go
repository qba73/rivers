package rivers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	levelSensor   = 1
	tempSensor    = 2
	voltageSensor = 3
)

// Sensor holds data representing
// a single sensor monted in a station.
type Sensor struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	ErrorCode int    `json:"err_code"`
}

// Station represents a station with
// multiple sensors.
type Station struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	RegionID   int      `json:"region_id"`
	RegionName string   `json:"region_name"`
	Lat        float64  `json:"lat"`
	Long       float64  `json:"long"`
	Sensors    []Sensor `json:"sensors"`
}

var validPeriod = map[string]bool{
	"day":   true,
	"week":  true,
	"month": true,
}

type errResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// LatestReading represents data received from a sensor.
type LatestReading struct {
	StationID   string
	StationName string
	SensorID    string
	RegionID    int
	Value       string
	Timestamp   string
	ErrCode     int
}

// Client holds data required to
// communicate with a web service.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient knows how to construct a new default rivers client.
// The client will be used to retrieve information about
// various measures recorded by sensors.
func NewClient() *Client {
	return &Client{
		BaseURL: "http://waterlevel.ie",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetLatest knows how to return latest readings from water level,
// water temperature and voltage level sensors installed in all
// stations in rivers in Ireland.
func (c *Client) GetLatest() ([]LatestReading, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geojson/latest", c.BaseURL), nil)
	if err != nil {
		return []LatestReading{}, err
	}

	var resp struct {
		Features []struct {
			Properties struct {
				StationRef  string `json:"station_ref"`
				StationName string `json:"station_name"`
				SensorRef   string `json:"sensor_ref"`
				RegionID    int    `json:"region_id"`
				Datetime    string `json:"datetime"`
				Value       string `json:"value"`
				ErrCode     int    `json:"err_code"`
			} `json:"properties"`
		}
	}

	if err := c.sendRequestJSON(req, &resp); err != nil {
		return []LatestReading{}, err
	}

	out := make([]LatestReading, len(resp.Features))

	for i, p := range resp.Features {
		reading := LatestReading{
			StationID:   p.Properties.StationRef,
			StationName: p.Properties.StationName,
			SensorID:    p.Properties.SensorRef,
			RegionID:    p.Properties.RegionID,
			Value:       p.Properties.Value,
			Timestamp:   p.Properties.Datetime,
			ErrCode:     p.Properties.ErrCode,
		}
		out[i] = reading
	}

	return out, nil
}

// GetDayLevel knows how to return water level readings recorded for
// last 24hr period for the given stationID number.
func (c *Client) GetDayLevel(stationID string) ([]SensorReading, error) {
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
func (c *Client) GetWeekLevel(stationID string) ([]SensorReading, error) {
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
func (c *Client) GetMonthLevel(stationID string) ([]SensorReading, error) {
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
func (c *Client) GetDayTemperature(stationID string) ([]SensorReading, error) {
	url, err := c.urlTemperature(stationID, "day")
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
func (c *Client) GetWeekTemperature(stationID string) ([]SensorReading, error) {
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
func (c *Client) GetMonthTemperature(stationID string) ([]SensorReading, error) {
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

// GetDayVoltage knows how to return sensor voltage data recorded over last 24h.
func (c *Client) GetDayVoltage(stationID string) ([]SensorReading, error) {
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
func (c *Client) urlLevel(stationID, period string) (string, error) {
	if !validPeriod[period] {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s", c.BaseURL, period, fileNameLevel(stationID)), nil
}

// urlTemperature takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *Client) urlTemperature(stationID, period string) (string, error) {
	if !validPeriod[period] {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s", c.BaseURL, period, fileNameTemperature(stationID)), nil
}

// urlVoltage takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *Client) urlVoltage(stationID, period string) (string, error) {
	if !validPeriod[period] {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s", c.BaseURL, period, fileNameVoltage(stationID)), nil
}

func (c *Client) sendRequestJSON(req *http.Request, v interface{}) error {
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

func (c *Client) sendRequestCSV(req *http.Request) ([]SensorReading, error) {
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

// GetLatest returns latests reading from all stations.
func GetLatest() ([]LatestReading, error) {
	c := NewClient()
	return c.GetLatest()
}
