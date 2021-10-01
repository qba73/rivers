package data

import (
	"encoding/json"
	"io"
)

type Station struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	River     string `json:"river"`
	UpdatedOn string `json:"updated_on"`
}

type Stations []*Station

func (s *Stations) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func GetStations() Stations {
	return stationList
}

var stationList = []*Station{
	{
		ID:        "001",
		Name:      "Station1",
		River:     "Shannon",
		UpdatedOn: "",
	},
	{
		ID:        "002",
		Name:      "Station2",
		River:     "Shanon",
		UpdatedOn: "",
	},
}
