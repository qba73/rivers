package rivers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type option func(*Puller) error

func WithSaver(st Store) option {
	return func(p *Puller) error {
		if st == nil {
			return errors.New("nil data saver")
		}
		p.Store = st
		return nil
	}
}

type Puller struct {
	Client   *Client
	Store    Store
	Interval time.Duration
}

func NewPuller(opts ...option) (*Puller, error) {
	riverClient := NewClient()
	store, err := NewSQLiteStore("data.db")
	if err != nil {
		return nil, fmt.Errorf("%w: creating data puller", err)
	}
	p := Puller{
		Client:   riverClient,
		Store:    store,
		Interval: 5 * time.Minute,
	}

	for _, opt := range opts {
		if err := opt(&p); err != nil {
			return nil, fmt.Errorf("creating data puller: %w", err)
		}
	}

	return &p, nil
}

func (p *Puller) Run() error {
	// Retrive new sensor data
	stationReadings, err := p.Client.GetLatestWaterLevels()
	if err != nil {
		return err
	}
	for _, reading := range stationReadings {
		err = p.Store.Save(reading)
		if err != nil {
			return err
		}
	}
	return nil
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
