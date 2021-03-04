// Package rivers provides functionality
// for manipulating time series data.
package river

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

// Level represents water level
// and
type Level struct {
	Timestamp time.Time
	Value     float64
}

// LoadCSV ...
func LoadCSV(filename string) ([]Level, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(data))
	// We are not interested in the CSV header, So, read it
	// before start looping through the lines.
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("error when reading csv file: %s", filename)
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var levels []Level

	for _, r := range records {
		tm := fixDate(r[0])
		t, err := time.Parse(time.RFC3339, tm)
		if err != nil {
			return nil, err
		}

		val, err := strconv.ParseFloat(r[1], 64)
		if err != nil {
			return nil, err
		}

		l := Level{
			Timestamp: t,
			Value:     val,
		}

		levels = append(levels, l)
	}
	return levels, nil
}

// fixDate is a helper function that make
// date format compatible with RFC3339.
func fixDate(s string) string {
	var date, time string
	fmt.Sscanf(s, "%s %s", &date, &time)
	return fmt.Sprintf("%sT%s:00Z", date, time)
}
