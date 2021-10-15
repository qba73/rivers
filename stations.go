package rivers

import (
	"encoding/json"
	"io"
	"os"
)

const (
	levelSensor   = 1
	tempSensor    = 2
	voltageSensor = 3
)

// Station is a measurement station
// situated in the lat long location.
type Station struct {
	Name  string
	RefNo string
	Lat   float64
	Long  float64
}

// LoadStations knows how to read stations json file.
// It takes a path to the json file and returns
// a struct Stations identified as geographical
// 'Features' in GeoJSON terms.
func LoadStations(name string) (StationsLatest, error) {
	f, err := os.Open(name)
	if err != nil {
		return StationsLatest{}, err
	}
	defer f.Close()
	return ReadStations(f)
}

// ReadStations knows how to unmarshal stations.
func ReadStations(r io.Reader) (StationsLatest, error) {
	var s StationsLatest

	if err := json.NewDecoder(r).Decode(&s); err != nil {
		return StationsLatest{}, err
	}

	return s, nil
}

// Stations represents station data.
type Stations struct {
	Type     string    `json:"type"`
	Crs      Crs       `json:"crs"`
	Features []Feature `json:"features,omitempty"`
}

// Crs represents CRS property.
type Crs struct {
	Type       string      `json:"type"`
	Properties CrsProperty `json:"properties"`
}

// CrsProperty is a property name.
type CrsProperty struct {
	Name string `json:"name"`
}

// Feature represents a feature with properties describing
// gauge name and gauge ID.
type Feature struct {
	Type       string   `json:"type"`
	Properties Property `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

// Property represents feature property.
type Property struct {
	Name string `json:"name"`
	Ref  string `json:"ref"`
}
