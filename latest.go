package rivers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// FeatureLatest represents a feature with properties describing
// latest readings from a sensor identified by ID.
type FeatureLatestGeo struct {
	Type       string            `json:"type"`
	Properties PropertyLatestGeo `json:"properties"`
	Geometry   GeometryGeo       `json:"geometry"`
}

// PropertyLatestGeo represents properties of a single gauge station.
type PropertyLatestGeo struct {
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

// GeometryGeo represents geometry coordinates of a gauge station.
type GeometryGeo struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Crs struct {
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
}

// StationsLatest represents geojson data describing
// all gauge stations and coordinates data.
type StationsLatestGeo struct {
	Type     string             `json:"type"`
	Crs      Crs                `json:"crs"`
	Features []FeatureLatestGeo `json:"features,omitempty"`
}

// ToJSON knows how to encode latest stations to json.
func (s StationsLatestGeo) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

// All knows how to return all gauge stations.
func (s StationsLatestGeo) GetAll() StationsLatestGeo {
	return s
}

// GetByName takes a feature (station) name and
// returns the Feature struct. If provided name is not
// found it returns an empty Feature.
func (s StationsLatestGeo) GetByName(name string) StationsLatestGeo {
	var features []FeatureLatestGeo

	for _, f := range s.Features {
		if f.Properties.StationName == name {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}

// GetByRefID takes station ID and returns
// matching features (stations). If Station Ref number does
// not exist it returns Stations struct with empty list of
// Features (stations/sensors).
func (s StationsLatestGeo) GetByID(ref string) StationsLatestGeo {
	var features []FeatureLatestGeo

	for _, f := range s.Features {
		if f.Properties.StationRef == ref {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}

// GetByStationAndSensorRef takes sensor ID and returns matching features
// (stations). If the sensor ID doesn't exist it returns Stations
// struct with an empty list of Fetures (stations/sensors).
func (s StationsLatestGeo) GetByStationAndSensorRef(station, sensor string) StationsLatestGeo {
	var features []FeatureLatestGeo

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
func (s StationsLatestGeo) GetByRegionID(regionID int) StationsLatestGeo {
	var features []FeatureLatestGeo

	for _, f := range s.Features {
		if f.Properties.RegionID == regionID {
			features = append(features, f)
		}
	}
	s.Features = features
	return s
}

// LoadStationsLatest knows how to read stations json file with
// the latest results. It takes a path to the json file and returns
// a struct Stations identified as geographical 'Features' in GeoJSON terms.
func LoadStationsLatest(name string) (StationsLatestGeo, error) {
	f, err := os.Open(name)
	if err != nil {
		return StationsLatestGeo{}, err
	}
	defer f.Close()
	return ReadStationsLatestFromJSON(f)
}

// ReadStationsLatestFromJSON knows how to unmarshal latest stations.
func ReadStationsLatestFromJSON(r io.Reader) (StationsLatestGeo, error) {
	var s StationsLatestGeo
	if err := json.NewDecoder(r).Decode(&s); err != nil {
		return StationsLatestGeo{}, err
	}
	return s, nil
}

// ===============================================================
// Station Handlers

type StationLatestStore interface {
	GetAll() StationsLatestGeo
	GetByID(id string) StationsLatestGeo
}

type JSONStore struct {
	Stations StationsLatestGeo
}

func NewJSONStore(pathname string) (StationsLatestGeo, error) {
	stations, err := LoadStationsLatest(pathname)
	if err != nil {
		return StationsLatestGeo{}, err
	}
	return stations, nil
}

type StationsHandler struct {
	l     *log.Logger
	Store StationLatestStore
}

func NewStationsHandler(l *log.Logger, store StationLatestStore) *StationsHandler {
	return &StationsHandler{l, store}
}

func (sh *StationsHandler) GetStations(w http.ResponseWriter, r *http.Request) {
	sx := sh.Store.GetAll()
	w.Header().Set("Content-Type", "application/json")
	if err := sx.ToJSON(w); err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (s *StationsHandler) GetStationsByID(w http.ResponseWriter, r *http.Request) {
	stations, err := LoadStationsLatest("testdata/latesttest.json")
	if err != nil {
		http.Error(w, "unable to load stations", http.StatusInternalServerError)
	}

	vars := mux.Vars(r)
	refid := vars["id"]

	sx := stations.GetByID(refid)
	w.Header().Set("Content-Type", "application/json")
	if err := sx.ToJSON(w); err != nil {
		http.Error(w, "unable to load stations", http.StatusInternalServerError)
		return
	}
}
