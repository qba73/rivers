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
	Features []Feature `json:"features"`
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

// FeatureLatest represents a feature with properties describing
// latest readings from a sensor identified by ID.
type FeatureLatest struct {
	Type       string         `json:"type"`
	Properties PropertyLatest `json:"properties"`
	Geometry   Geometry       `json:"geometry"`
}

// PropertyLatest represents properties of a single gauge station.
type PropertyLatest struct {
	StationRef  string `json:"station_ref"`
	StationName string `json:"station_name"`
	SensorRef   string `json:"sensor_ref"`
	RegionID    int    `json:"region_id"`
	Timestamp   string `json:"datetime"`
	Value       string `json:"value"`
	ErrCode     int    `json:"err_code"`
	URL         string `json:"url"`
	CSVFile     string `json:"csv_file"`
}

// Geometry represents geometry coordinates of a gauge station.
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// StationsLatest represents geojson data describing
// all gauge stations and coordinates data.
type StationsLatest struct {
	Type     string          `json:"type"`
	Crs      Crs             `json:"crs"`
	Features []FeatureLatest `json:"features"`
}

// All knows how to return all gauge stations.
func (s StationsLatest) All() StationsLatest {
	return s
}

// ByName takes a feature (station) name and
// returns the Feature struct. If provided name is not
// found it returns an empty Feature.
func (s StationsLatest) ByName(name string) StationsLatest {
	var features []FeatureLatest

	for _, f := range s.Features {
		if f.Properties.StationName == name {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}

// ByRefID takes station ID and returns
// matching features (stations). If Station Ref number does
// not exist it returns Stations struct with empty list of
// Features (stations/sensors).
func (s StationsLatest) ByRefID(ref string) StationsLatest {
	var features []FeatureLatest

	for _, f := range s.Features {
		if f.Properties.StationRef == ref {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}

// ByStationAndSensorRef takes sensor ID and returns matching features
// (stations). If the sensor ID doesn't exist it returns Stations
// struct with an empty list of Fetures (stations/sensors).
func (s StationsLatest) ByStationAndSensorRef(station, sensor string) StationsLatest {
	var features []FeatureLatest

	for _, f := range s.Features {
		if f.Properties.StationRef == station && f.Properties.SensorRef == sensor {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}

// ByRegionID knows how to return stations assigned
// to the given region identified by regionID.
func (s StationsLatest) ByRegionID(regionID int) StationsLatest {
	var features []FeatureLatest

	for _, f := range s.Features {
		if f.Properties.RegionID == regionID {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}
