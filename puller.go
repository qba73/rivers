package rivers

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Puller struct {
	Client   *Client
	Store    Saver
	Interval time.Duration
}

func NewPuller() (*Puller, error) {
	riverClient := NewClient()
	store, err := NewFileStore("data.txt")
	if err != nil {
		return nil, fmt.Errorf("%w: creating data puller", err)
	}
	p := Puller{
		Client:   riverClient,
		Store:    store,
		Interval: 5 * time.Minute,
	}
	return &p, nil
}

func (p *Puller) Run() error {
	// Retrive new sensor data
	stationReadings, err := p.Client.GetLatestWaterLevels()
	if err != nil {
		return err
	}
	return p.Store.Save(stationReadings)
}

// RunPuller holds all required machinery
// to run the river data puller.
func RunPuller() {
	log := log.New(os.Stdout, "PULLER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	p, err := NewPuller()
	if err != nil {
		panic(err)
	}
	duration, err := time.ParseDuration("1m")
	if err != nil {
		panic(err)
	}

	log.Println("rivers: Start pulling sensor data")

	for range time.NewTicker(p.Interval).C {
		log.Println("rivers: Pulling data")
		if err := p.Run(); err != nil {
			break
		}
		log.Printf("rivers: Data saved. Resuming in %s", duration)
	}
}
