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
type WaterLevelReading struct {
	Name      string
	RefID     string
	Timestamp time.Time
	Value     int
}

type WaterTemperatureReading struct {
	Name      string
	RefID     string
	Timestamp time.Time
	Value     float64
}

type VoltageReading struct {
	Name      string
	RefID     string
	Timestamp time.Time
	Value     float64
}

// String prints out information about the reading.
func (rd WaterLevelReading) String() string {
	return fmt.Sprintf("Name: %s, Time: %s, Value: %v", rd.Name, rd.Timestamp, rd.Value)
}

// LoadWaterLevelCSV knows how to open and read given csv file.
// Upon successful run it returns a slice of level structs.
func LoadWaterLevelCSV(path string) ([]WaterLevelReading, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadWaterLevelCSV(f)
}

// ReadWaterLevelCSV knows how to read a csv file containing
// readings from a gauge in a format: timestamp,level
func ReadWaterLevelCSV(r io.Reader) ([]WaterLevelReading, error) {
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

	levels := make([]WaterLevelReading, len(records))
	for i, r := range records {
		level, err := processWaterLevelRecord(r)
		if err != nil {
			return nil, fmt.Errorf("error processing csv record: %v", err)
		}
		levels[i] = level
	}
	return levels, nil
}

// ReadWaterTemperatureCSV knows how to read a csv file containing
// readings from a gauge in a format: `timestamp,value`. The value
// represents temp in Celsius.
func ReadWaterTemperatureCSV(r io.Reader) ([]WaterTemperatureReading, error) {
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

	levels := make([]WaterTemperatureReading, len(records))
	for i, r := range records {
		level, err := processWaterTempRecord(r)
		if err != nil {
			return nil, fmt.Errorf("error processing csv record: %v", err)
		}
		levels[i] = level
	}
	return levels, nil
}

func processWaterLevelRecord(r []string) (WaterLevelReading, error) {
	tm, err := processTimestamp(r)
	if err != nil {
		return WaterLevelReading{}, err
	}
	val, err := processWaterLevelValue(r)
	if err != nil {
		return WaterLevelReading{}, err
	}
	return WaterLevelReading{Timestamp: tm, Value: val}, nil
}

func processWaterTempRecord(r []string) (WaterTemperatureReading, error) {
	tm, err := processTimestamp(r)
	if err != nil {
		return WaterTemperatureReading{}, err
	}
	val, err := processWaterTempValue(r)
	if err != nil {
		return WaterTemperatureReading{}, err
	}
	return WaterTemperatureReading{Timestamp: tm, Value: val}, nil
}

func processTimestamp(record []string) (time.Time, error) {
	if len(record) < 1 {
		return time.Time{}, fmt.Errorf("processing timestamp: invalid record %v", record)
	}
	return time.Parse(gaugeTimeFormat, record[0])
}

func processWaterLevelValue(record []string) (int, error) {
	if len(record) != 2 {
		return 0, fmt.Errorf("processing water level value: invalid record %v", record)
	}
	return toMillimeters(record[1])
}

func processWaterTempValue(record []string) (float64, error) {
	if len(record) != 2 {
		return 0, fmt.Errorf("processing water temp value: invalid record %v", record)
	}
	return strconv.ParseFloat(record[1], 64)
}

func ReadGroupCSV(r io.Reader) ([]WaterLevelReading, error) {
	csvreader := csv.NewReader(r)
	records, err := csvreader.ReadAll()
	if err != nil {
		return nil, err
	}
	return parseStationGroup(records)
}

func parseStationGroup(records [][]string) ([]WaterLevelReading, error) {
	if len(records) < 2 {
		return nil, fmt.Errorf("parsing station groups: empty records %v", records)
	}
	if len(records[0]) < 2 {
		return nil, errors.New("parsing station groups: missing station")
	}
	var gsr []WaterLevelReading
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
			levelValue, err := toMillimeters(reading)
			if err != nil {
				return nil, fmt.Errorf("parsing station groups: %w", err)
			}
			gr := WaterLevelReading{
				// Some headers in csv files come with empty spaces, so trim them.
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
	StationID  int       `json:"station_id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Readtime   time.Time `json:"readtime"`
	WaterLevel int       `json:"water_level"`
}

type StationGroupReading struct {
	GroupID      int       `json:"group_id"`
	GroupName    string    `json:"group_name"`
	StationID    string    `json:"station_id"`
	Name         string    `json:"name,omitempty"`
	Readtime     time.Time `json:"readtime"`
	ReadingValue int       `json:"reading_value"`
}
