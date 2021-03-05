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

// ReadStations ...
func ReadStations(r io.Reader) (Stations, error) {
	var s Stations
	if err := Unmarshal(r, &s); err != nil {
		return Stations{}, err
	}
	return s, nil
}

// Feature represents a single gauge.
type Feature struct {
	Type       string   `json:"type"`
	Properties Property `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

// Property represents properties of agauge station.
type Property struct {
	Name string `json:"name"`
	Ref  string `json:"ref"`
}

// Geometry represents geometry coordinates of a gauge station.
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Station represent a measurement station with installed gauge.
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

// GetAllFeatures ...
func (s Stations) GetAllFeatures() []Feature {
	return s.Features
}

// GetFeatureByName takes a feature (station) name and
// returns the Feature struct. If provided name is not
// found it returns an empty Feature.
func (s Stations) GetFeatureByName(name string) Feature {
	for _, f := range s.Features {
		if f.Properties.Name == name {
			return f
		}
	}
	return Feature{}
}

// GetFeatureByRef takes feature reference number and returns
// matching feature. Othervise it returns an empty Feature struct.
func (s Stations) GetFeatureByRef(ref string) Feature {
	for _, f := range s.Features {
		if f.Properties.Ref == ref {
			return f
		}
	}
	return Feature{}
}

// GetAllStations ...
func (s Stations) GetAllStations() ([]Station, error) {
	var stations []Station
	return stations, nil
}
