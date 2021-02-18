package rivers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Stations represents geojson data describing
// all gauge stations and coordinates data.
type Stations struct {
	Type string `json:"type"`
	Crs  struct {
		Type       string
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
	} `json:"crs"`
	Features []Feature `json:"features"`
}

// Feature represents a single gauge.
type Feature struct {
	Type       string `json:"type"`
	Properties struct {
		Name string `json:"name"`
		Ref  string `json:"ref"`
	} `json:"properties"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

// Station represents a gauge on the river identified
// by name and a reference number.
type Station struct {
	Type       string `json:"type"`
	Properties struct {
		Name string `json:"type"`
		Ref  string `json:"ref"`
	} `json:"properties"`
}

// LoadStations knows how to read stations json file.
// It returns a struct Stations with gauge stations known as
// geographical 'Features' (GIS terminology).
func LoadStations(filename string) (Stations, error) {
	var stations Stations
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Stations{}, fmt.Errorf("error reading file: %s", filename)
	}

	if err := json.Unmarshal(data, &stations); err != nil {
		return Stations{}, fmt.Errorf("error unmarshalling stations")
	}
	return stations, nil
}
