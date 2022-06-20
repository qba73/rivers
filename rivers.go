package rivers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// gaugeTimeFormat represents formatted time field in csv data files.
	gaugeTimeFormat = "2006-01-02 15:04"
)

// Reading represents water level recorded
// by a gauge at the particular time.
type Reading struct {
	Name      string
	RefID     string
	Timestamp time.Time
	Value     float64
}

// String prints out information about the reading.
func (rd Reading) String() string {
	return fmt.Sprintf("Name: %s, Time: %s, Value: %v", rd.Name, rd.Timestamp, rd.Value)
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
func ReadCSV(r io.Reader) ([]Reading, error) {
	var levels []Reading

	csvreader := csv.NewReader(r)
	// We are not interested in the CSV header. We read it
	// before start looping through the lines (records).
	if _, err := csvreader.Read(); err != nil {
		return nil, errors.New("error when reading csv file")
	}

	records, err := csvreader.ReadAll()
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

func ReadGroupCSV(r io.Reader) ([]Reading, error) {
	csvreader := csv.NewReader(r)
	records, err := csvreader.ReadAll()
	if err != nil {
		return nil, err
	}
	return parseStationGroup(records)
}

func parseStationGroup(records [][]string) ([]Reading, error) {
	if len(records) < 2 {
		return nil, fmt.Errorf("empty records")
	}
	if len(records[0]) < 2 {
		return nil, fmt.Errorf("missing station")
	}
	var gsr []Reading
	stationNames := records[0][1:]

	for _, record := range records[1:] {
		timestamp, err := time.Parse(gaugeTimeFormat, record[0])
		if err != nil {
			return nil, err
		}
		for i, reading := range record[1:] {
			// If reading value from sensor is not yet present we
			// do not attemp to process it.
			if reading == "" {
				continue
			}
			levelValue, err := strconv.ParseFloat(reading, 64)
			if err != nil {
				return nil, err
			}
			gr := Reading{
				// Some headers in csv files come with empty spaces. We trim them.
				Name:      strings.TrimSpace(stationNames[i]),
				Timestamp: timestamp,
				Value:     levelValue,
			}
			gsr = append(gsr, gr)
		}
	}
	return gsr, nil
}

// WaterLevelProvider is the interface that wraps the GetLatestWaterLevels method.
//
// GetLatestWaterLevels returns a slice of water level readings from sensors.
type WaterLevelProvider interface {
	GetLatestWaterLevels() []StationWaterLevelReading
}

// StationWaterLevelReading represents data received
// from a water level sensor.
type StationWaterLevelReading struct {
	StationID  string    `json:"station_id"`
	Name       string    `json:"name,omitempty"`
	RegionID   int       `json:"region_id,omitempty"`
	Readtime   time.Time `json:"readtime"`
	WaterLevel float64   `json:"water_level"`
}
