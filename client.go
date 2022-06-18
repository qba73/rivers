package rivers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

const (
	libVersion = "0.0.1"

	// Sensor types
	levelSensor   = 1
	tempSensor    = 2
	voltageSensor = 3

	sensorTypeLevel = "0001"
	sensorTypeTemp  = "0002"
)

type response struct {
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

// Sensor holds data from a station.
type Sensor struct {
	StationID   string `json:"station_id"`
	StationName string `json:"station_name"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Timestamp   string `json:"timestamp"`
	ErrorCode   int    `json:"err_code"`
	RegionID    string `json:"region_id"`
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

type errResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SensorReading represents data received from a sensor.
type SensorReading struct {
	StationID   string
	StationName string
	SensorID    string
	RegionID    int
	Value       float64
	Timestamp   time.Time
	ErrCode     int
}

// Client holds data required to communicate with the web service.
type Client struct {
	UserAgent  string
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient knows how to construct a new default rivers client.
// The client will be used to retrieve information about
// various measures recorded by sensors.
func NewClient() *Client {
	return &Client{
		UserAgent: "Rivers/" + libVersion,
		BaseURL:   "http://waterlevel.ie",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetLatestWaterLevels returns latest water level readings from sensors.
func (c *Client) GetLatestWaterLevels() ([]StationWaterLevelReading, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/geojson/latest", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}
	var res response
	if err := c.sendRequestJSON(req, &res); err != nil {
		return nil, err
	}
	var out []StationWaterLevelReading
	for _, p := range res.Features {
		if p.Properties.SensorRef != sensorTypeLevel {
			continue
		}
		t, err := time.Parse(time.RFC3339, p.Properties.Datetime)
		if err != nil {
			return nil, err
		}
		wl, err := strconv.ParseFloat(p.Properties.Value, 64)
		if err != nil {
			return nil, err
		}
		reading := StationWaterLevelReading{
			StationID:  p.Properties.StationRef,
			Name:       p.Properties.StationName,
			RegionID:   p.Properties.RegionID,
			Readtime:   t,
			WaterLevel: wl,
		}
		out = append(out, reading)
	}
	return out, nil
}

// GetDayLevel knows how to return water level readings recorded for
// last 24hr period for the given stationID number.
func (c *Client) GetDayLevel(stationID string) ([]Reading, error) {
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
func (c *Client) GetWeekLevel(stationID string) ([]Reading, error) {
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
func (c *Client) GetMonthLevel(stationID string) ([]Reading, error) {
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
func (c *Client) GetDayTemperature(stationID string) ([]Reading, error) {
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
func (c *Client) GetWeekTemperature(stationID string) ([]Reading, error) {
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
func (c *Client) GetMonthTemperature(stationID string) ([]Reading, error) {
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
func (c *Client) GetDayVoltage(stationID string) ([]Reading, error) {
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

// GetGroupWaterLevel returns water level readings for
// stations that belong to provided groupID.
// The value of roupID should be between 1 and 28.
func (c *Client) GetGroupWaterLevel(groupID int) ([]Reading, error) {
	if groupID < 1 || groupID > 28 {
		return nil, fmt.Errorf("invalid groupID %d, expecting value between 1 and 28", groupID)
	}
	url := fmt.Sprintf("%s/data/group/group_%d.csv", c.BaseURL, groupID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.sendStationGroupRequestCSV(req)
}

var validPeriods = []string{"day", "week", "month"}

// urlLevel takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *Client) urlLevel(stationID, period string) (string, error) {
	if !slices.Contains(validPeriods, period) {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s_000%v.csv", c.BaseURL, period, stationID, levelSensor), nil
}

// urlTemperature takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *Client) urlTemperature(stationID, period string) (string, error) {
	if !slices.Contains(validPeriods, period) {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s_000%v.csv", c.BaseURL, period, stationID, tempSensor), nil
}

// urlVoltage takes stationid and time period and builds a valid url.
// If the period is not valid it errors. Period value should be
// one of 'day', 'week' or 'month'.
func (c *Client) urlVoltage(stationID, period string) (string, error) {
	if !slices.Contains(validPeriods, period) {
		return "", fmt.Errorf("invalid period %q, expecting one of 'day', 'week', 'month'", period)
	}
	return fmt.Sprintf("%s/data/%s/%s_000%v.csv", c.BaseURL, period, stationID, voltageSensor), nil
}

func (c *Client) sendRequestJSON(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", c.UserAgent)
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
	return json.NewDecoder(res.Body).Decode(&v)
}

func (c *Client) sendRequestCSV(req *http.Request) ([]Reading, error) {
	req.Header.Set("Content-Type", "text/csv")
	req.Header.Set("Accept", "text/csv")
	req.Header.Set("User-Agent", c.UserAgent)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	return ReadCSV(res.Body)
}

func (c *Client) sendStationGroupRequestCSV(req *http.Request) ([]Reading, error) {
	req.Header.Set("Content-Type", "text/csv")
	req.Header.Set("Accept", "text/csv")
	req.Header.Set("User-Agent", c.UserAgent)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	return ReadGroupCSV(res.Body)
}

// GetLatestWaterLevels returns latests readings from all stations.
//
// This func uses default rivers' client under the hood.
func GetLatestWaterLevels() ([]StationWaterLevelReading, error) {
	return NewClient().GetLatestWaterLevels()
}
