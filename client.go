package rivers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
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

type errResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Sensor represents a sensor installed in a station.
type Sensor struct {
	StationID   string `json:"station_id"`
	StationName string `json:"station_name"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Timestamp   string `json:"timestamp"`
	ErrorCode   int    `json:"err_code"`
	RegionID    string `json:"region_id"`
}

// Station represents a station with multiple sensors.
type Station struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	RegionID   int      `json:"region_id"`
	RegionName string   `json:"region_name"`
	Lat        float64  `json:"lat"`
	Long       float64  `json:"long"`
	Sensors    []Sensor `json:"sensors"`
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
func (c *Client) GetLatestWaterLevels(ctx context.Context) ([]StationWaterLevelReading, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/geojson/latest", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}
	var res response
	if err := c.sendRequestJSON(req, &res); err != nil {
		return nil, err
	}

	var readings []StationWaterLevelReading
	for _, p := range res.Features {
		if p.Properties.SensorRef != sensorTypeLevel {
			continue
		}

		stationID, err := fromStrToInt(p.Properties.StationRef)
		if err != nil {
			return nil, fmt.Errorf("converting station id from string to int: %w", err)
		}

		t, err := time.Parse(time.RFC3339, p.Properties.Datetime)
		if err != nil {
			return nil, fmt.Errorf("parsing reading time: %w", err)
		}

		wl, err := toMillimeters(p.Properties.Value)
		if err != nil {
			return nil, err
		}
		reading := StationWaterLevelReading{
			StationID:  stationID,
			Name:       p.Properties.StationName,
			Readtime:   t,
			WaterLevel: wl,
		}
		readings = append(readings, reading)
	}
	return readings, nil
}

// GetDayLevel knows how to return water level readings recorded for
// last 24hr period for the given stationID number.
func (c *Client) GetDayLevel(ctx context.Context, stationID string) ([]WaterLevelReading, error) {
	url, err := c.urlLevel(stationID, "day")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.requestWaterLevelCSV(req)
}

// GetWeekLevel knows how to return water level readings recorded for
// last week period for the given stationID number.
func (c *Client) GetWeekLevel(ctx context.Context, stationID string) ([]WaterLevelReading, error) {
	url, err := c.urlLevel(stationID, "week")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.requestWaterLevelCSV(req)
}

// GetMonthLevel knows how to return water level readings recorded for
// last 4 weeks period for the given stationID number.
func (c *Client) GetMonthLevel(ctx context.Context, stationID string) ([]WaterLevelReading, error) {
	url, err := c.urlLevel(stationID, "month")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.requestWaterLevelCSV(req)
}

// GetDayTemperature knows how to return water temperature
// recorded for last 24hr period for the given stationID number.
func (c *Client) GetDayTemperature(ctx context.Context, stationID string) ([]WaterTemperatureReading, error) {
	url, err := c.urlTemperature(stationID, "day")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.requestWaterTemperatureCSV(req)
}

// GetWeekTemperature knows how to return water temperature
// recorded for last week period for the given stationID number.
func (c *Client) GetWeekTemperature(ctx context.Context, stationID string) ([]WaterTemperatureReading, error) {
	url, err := c.urlTemperature(stationID, "week")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.requestWaterTemperatureCSV(req)
}

// GetMonthTemperature knows how to return water temperature
// recorded for last 4 weeks period for the given stationID number.
func (c *Client) GetMonthTemperature(ctx context.Context, stationID string) ([]WaterTemperatureReading, error) {
	url, err := c.urlTemperature(stationID, "month")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.requestWaterTemperatureCSV(req)
}

// GetGroupWaterLevel returns water level readings for
// stations that belong to the given groupID.
//
// The value of roupID should be between 1 and 28.
func (c *Client) GetGroupWaterLevel(ctx context.Context, groupID int) ([]StationWaterLevelReading, error) {
	if groupID < 1 || groupID > 28 {
		return nil, fmt.Errorf("invalid groupID %d, expecting value between 1 and 28", groupID)
	}
	url := fmt.Sprintf("%s/data/group/group_%d.csv", c.BaseURL, groupID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	groupReadings, err := c.sendStationGroupRequestCSV(req)
	if err != nil {
		return nil, err
	}

	var readings []StationWaterLevelReading
	for _, reading := range groupReadings {
		station := StationWaterLevelReading{
			Name:       reading.Name,
			Readtime:   reading.Timestamp,
			WaterLevel: reading.Value,
		}
		readings = append(readings, station)
	}
	return readings, nil
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

func (c *Client) sendRequestJSON(req *http.Request, v any) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", c.UserAgent)

	res, err := c.sendRequestWithBackoff(req)
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

func (c *Client) sendRequestWithBackoff(req *http.Request) (*http.Response, error) {
	res, err := c.HTTPClient.Do(req)
	base := time.Second
	cap := time.Minute
	for backoff := base; err != nil; backoff <<= 1 {
		if backoff > cap {
			backoff = cap
		}
		jitter := rand.Int63n(int64(backoff * 3))
		time.Sleep(base + time.Duration(jitter))
		res, err = c.HTTPClient.Do(req)
	}
	return res, err
}

func (c *Client) requestWaterLevelCSV(req *http.Request) ([]WaterLevelReading, error) {
	req.Header.Set("Content-Type", "text/csv")
	req.Header.Set("Accept", "text/csv")
	req.Header.Set("User-Agent", c.UserAgent)

	res, err := c.sendRequestWithBackoff(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := checkResponseStatusCode(res); err != nil {
		return nil, err
	}
	return ReadWaterLevelCSV(res.Body)
}

func (c *Client) requestWaterTemperatureCSV(req *http.Request) ([]WaterTemperatureReading, error) {
	req.Header.Set("Content-Type", "text/csv")
	req.Header.Set("Accept", "text/csv")
	req.Header.Set("User-Agent", c.UserAgent)

	res, err := c.sendRequestWithBackoff(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := checkResponseStatusCode(res); err != nil {
		return nil, err
	}
	return ReadWaterTemperatureCSV(res.Body)
}

func (c *Client) sendStationGroupRequestCSV(req *http.Request) ([]WaterLevelReading, error) {
	req.Header.Set("Content-Type", "text/csv")
	req.Header.Set("Accept", "text/csv")
	req.Header.Set("User-Agent", c.UserAgent)

	res, err := c.sendRequestWithBackoff(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := checkResponseStatusCode(res); err != nil {
		return nil, err
	}
	return ReadGroupCSV(res.Body)
}

// checkResponseStatusCode takes http Respopnse and validates
// HTTP response status code.
// It errors if the status code is not 2xx or 3xx.
func checkResponseStatusCode(res *http.Response) error {
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	return nil
}

// toMillimeters is a helper func that takes a string representing a float value
// of water level sensor reading in meters and converts it to integer representing
// water level in millimeters.
// It errors if the input string does not represent float value.
func toMillimeters(s string) (int, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("convering to millimeters: %w", err)
	}
	return int(v * 1000), nil
}

// fromStrToInt is a helper func that takes a string representing
// stationID and returns stationID as int (with trimmed leading zeros).
func fromStrToInt(s string) (int, error) {
	st := strings.TrimLeft(s, "0")
	return strconv.Atoi(st)
}

func writeReadingsTo(w io.Writer, readings []StationWaterLevelReading) {
	for _, reading := range readings {
		fmt.Fprintf(w, "time: %s, station: %s, id: %d, level: %d\n",
			reading.Readtime, reading.Name, reading.StationID, reading.WaterLevel)
	}
}

// GetLatestWaterLevels returns latests readings from all stations.
//
// This func uses default rivers' client under the hood.
func GetLatestWaterLevels(ctx context.Context) ([]StationWaterLevelReading, error) {
	return NewClient().GetLatestWaterLevels(ctx)
}

// RunCLI executes program and prints out latest recorded water levels.
func RunCLI() {
	ctx, shutdown := signal.NotifyContext(context.Background(), os.Interrupt)
	defer shutdown()
	readings, err := GetLatestWaterLevels(ctx)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	writeReadingsTo(os.Stdout, readings)
	os.Exit(0)
}
