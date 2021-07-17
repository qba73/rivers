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
	// GaugeTimeFormat represents formatted time field in csv data files.
	GaugeTimeFormat = "2006-01-02 15:04"
)

// Level represents water level recorded
// by a gauge at the particular time.
type SensorReading struct {
	Timestamp time.Time
	Value     float64
}

// LoadCSV knows how to open and read given csv file.
// Upon successful run it returns a slice of level structs.
func LoadCSV(path string) ([]SensorReading, error) {
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
func ReadCSV(csvFile io.Reader) ([]SensorReading, error) {
	var levels []SensorReading

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

func processRecord(r []string) (SensorReading, error) {
	tm, err := processTimestamp(r)
	if err != nil {
		return SensorReading{}, err
	}
	val, err := processValue(r)
	if err != nil {
		return SensorReading{}, err
	}

	return SensorReading{Timestamp: tm, Value: val}, nil
}

func processTimestamp(record []string) (time.Time, error) {
	tm, err := time.Parse(GaugeTimeFormat, record[0])
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
