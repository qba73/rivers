package river

import (
	"encoding/json"
	"io"
	"os"
)

// LoadStations knows how to read stations json file.
// It takes a path to the json file and returns
// a struct Stations identified as geographical
// 'Features' in GeoJSON terms.
func LoadStations(name string) (Stations, error) {
	f, err := os.Open(name)
	if err != nil {
		return Stations{}, err
	}
	defer f.Close()
	return ReadStations(f)
}

// ReadStations knows how to unmarshal stations.
func ReadStations(r io.Reader) (Stations, error) {
	var s Stations

	if err := json.NewDecoder(r).Decode(&s); err != nil {
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

// Property represents properties of a single
// gauge station.
type Property struct {
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

// All ...
func (s Stations) All() Stations {
	return s
}

// ByName takes a feature (station) name and
// returns the Feature struct. If provided name is not
// found it returns an empty Feature.
func (s Stations) ByName(name string) Stations {
	var features []Feature

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
func (s Stations) ByRefID(ref string) Stations {
	var features []Feature

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
func (s Stations) ByStationAndSensorRef(station, sensor string) Stations {
	var features []Feature
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
func (s Stations) ByRegionID(regionID int) Stations {
	var features []Feature
	for _, f := range s.Features {
		if f.Properties.RegionID == regionID {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}
