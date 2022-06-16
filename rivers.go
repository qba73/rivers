package rivers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	// gaugeTimeFormat represents formatted time field in csv data files.
	gaugeTimeFormat = "2006-01-02 15:04"
)

// Reading represents water level recorded
// by a gauge at the particular time.
type Reading struct {
	Timestamp time.Time
	Value     float64
}

// LoadCSV knows how to open and read given csv file.
// Upon successful run it returns a slice of level structs.
func LoadCSV(path string) ([]Reading, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadCSV(f)
}

// ReadCSV knows how to read a csv file containing
// readings from a gauge in a format:
// timestamp,level
func ReadCSV(csvFile io.Reader) ([]Reading, error) {
	var levels []Reading

	r := csv.NewReader(csvFile)
	// We are not interested in the CSV header. We read it
	// before start looping through the lines (records).
	if _, err := r.Read(); err != nil {
		return nil, errors.New("error when reading csv file")
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, r := range records {
		level, err := processRecord(r)
		if err != nil {
			return nil, fmt.Errorf("error processing csv record: %v", err)
		}
		levels = append(levels, level)
	}
	return levels, nil
}

func processRecord(r []string) (Reading, error) {
	tm, err := processTimestamp(r)
	if err != nil {
		return Reading{}, err
	}
	val, err := processValue(r)
	if err != nil {
		return Reading{}, err
	}

	return Reading{Timestamp: tm, Value: val}, nil
}

func processTimestamp(record []string) (time.Time, error) {
	tm, err := time.Parse(gaugeTimeFormat, record[0])
	if err != nil {
		return time.Time{}, err
	}
	return tm, nil
}

func processValue(record []string) (float64, error) {
	val, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return 0, nil
	}
	return val, nil
}

// WaterLevelProvider is the interface that wraps the GetLatestWaterLevels method.
//
// GetLatestWaterLevels returns a slice of water level readings from sensors.
type WaterLevelProvider interface {
	GetLatestWaterLevels() []StationWaterLevelReading
}

// StationWaterLevelReading stores data obtained from
// the water level sensor.
type StationWaterLevelReading struct {
	StationID  string    `json:"station_id"`
	Readtime   time.Time `json:"readtime"`
	WaterLevel float64   `json:"water_level"`
}

// RunServer holds all required machinery
// to run the river web server.
func RunServer() {
	fmt.Println("Running rivers server...")
}

// RunPuller holds all required machinery
// to run the river data puller.
func RunPuller() {
	fmt.Println("Running the water data puller...")
}
