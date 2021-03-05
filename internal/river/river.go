package river

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// Level represents water level recorded
// by a gauge at the particular time.
type Level struct {
	Timestamp time.Time
	Value     float64
}

// LoadCSV knows how to open and read given csv file.
// The csv file
func LoadCSV(path string) ([]Level, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ReadCSV(f)
}

// ReadCSV knows how to read a csv file containing
// readings from a gauge in a format:
// timestamp,level
func ReadCSV(csvFile io.Reader) ([]Level, error) {
	var levels []Level

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

func processRecord(r []string) (Level, error) {
	tm, err := processTimestamp(r)
	if err != nil {
		return Level{}, err
	}
	val, err := processValue(r)
	if err != nil {
		return Level{}, err
	}

	return Level{Timestamp: tm, Value: val}, nil
}

func processTimestamp(record []string) (time.Time, error) {
	ft := fixTimestamp(record[0])
	tm, err := time.Parse(time.RFC3339, ft)
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

// fixTimeStampo is a helper function that makes date field
// compliant with RFC3339.
func fixTimestamp(s string) string {
	var date, time string
	fmt.Sscanf(s, "%s %s", &date, &time)
	return fmt.Sprintf("%sT%s:00Z", date, time)
}
