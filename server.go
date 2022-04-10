package rivers

import (
	"fmt"
	"time"
)

type WaterLevelProvider interface {
	GetLatestReading() []StationTempReading
}

type StationTempReading struct {
	StationID   string
	Readtime    time.Time
	Temperature float64
}

type StationWaterLevelReading struct {
	StationID  string
	Readtime   time.Time
	WaterLevel float64
}

type Server struct {
	DataProvider WaterLevelProvider
}

func NewServer(wlp WaterLevelProvider) *Server {
	return &Server{
		DataProvider: wlp,
	}
}

// RunServer holds all required machinery
// to run the river web server.
func RunServer() {
	fmt.Println("Staring server...")
}
