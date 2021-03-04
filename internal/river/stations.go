package river

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

var mu sync.Mutex

// Unmarshal knows how to unmarshal data from the reader
// into the specified value.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// LoadStations knows how to read stations json file.
// It returns a struct Stations with gauge stations known as
// geographical 'Features' (GIS terminology).
func LoadStations(path string, v interface{}) error {
	mu.Lock()
	defer mu.Unlock()

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return Unmarshal(f, v)
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

type Station struct {
	RefNo       string `json:"station_ref"`
	Name        string `json:"station_name"`
	SensorRefNo string `json:"sensor_ref"`
	RegionID    int    `json:"region_id"`
}

// Stations represents geojson data describing
// all gauge stations and coordinates data.
type Stations struct {
	Type string `json:"type"`
	Crs  struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
	} `json:"crs"`
	Features []Feature `json:"features"`
}

// GetAll ...
func (s Stations) GetAll() ([]Feature, error) {
	return []Feature{}, nil
}

// GetByName
func (s Stations) GetByName(n string) ([]Feature, error) {
	return []Feature{}, nil
}

// GetById

// GetByGroup
