package rivers

import (
	"time"
)

const (
	levelSensor   = 1
	tempSensor    = 2
	voltageSensor = 3
)

type Sensor struct {
	RefID     string `json:"id"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	ErrorCode int    `json:"err_code"`
}

// Station is a measurement station
// situated in a location identified by
// lat and long coordinates.
type Station struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	RegionID   int       `json:"region_id"`
	RegionName string    `json:"region_name"`
	Datetime   time.Time `json:"datetime"`
	Lat        float64   `json:"lat"`
	Long       float64   `json:"long"`
	Sensors    []Sensor  `json:"sensors"`
}

// Stations represent all stations installed in rivers.
type Stations []Station

func NewStations() (Stations, error) {
	//stations, err := LoadStationsLatest("testdata/latestjson")
	return Stations{}, nil
}

func (s Stations) GetAll() Stations {
	return Stations{}
}

func (s Stations) GetByName(name string) Stations {
	return Stations{}
}

func (s Stations) GetByRefID(refid string) Stations {
	return Stations{}
}

func (s Stations) GetByRegionID(regionID int) Stations {
	return Stations{}
}

func (s Stations) GetByRegionName(region string) Stations {
	return Stations{}
}

func GetAllStations() (Stations, error) {
	stations, err := NewStations()
	if err != nil {
		return Stations{}, err
	}
	return stations, nil
}
