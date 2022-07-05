package rivers

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type option func(*Puller) error

func WithInterval(duration string) option {
	return func(p *Puller) error {
		if duration == "" {
			return errors.New("setting up pulling interval: nil duration")
		}
		duration, err := time.ParseDuration(duration)
		if err != nil {
			return fmt.Errorf("setting up pulling interval: %w", err)
		}
		p.Interval = duration
		return nil
	}
}

type Puller struct {
	Client      *Client
	ReadingRepo *ReadingsRepo
	Interval    time.Duration
}

func NewPuller(opts ...option) (*Puller, error) {
	client := NewClient()
	store, err := NewSQLiteStore("waterlevels.db")
	if err != nil {
		return nil, fmt.Errorf("%w: creating data puller", err)
	}
	repo := OpenReadingsRepo(store)

	p := Puller{
		Client:      client,
		ReadingRepo: repo,
		Interval:    5 * time.Minute,
	}

	for _, opt := range opts {
		if err := opt(&p); err != nil {
			return nil, fmt.Errorf("creating data puller: %w", err)
		}
	}

	return &p, nil
}

func (p *Puller) Run() error {
	stationReadings, err := p.Client.GetLatestWaterLevels()
	if err != nil {
		return err
	}
	for _, reading := range stationReadings {
		err = p.ReadingRepo.Add(reading)
		if err != nil {
			return err
		}
	}
	return nil
}

// RunPuller holds all required machinery
// to run the river data puller.
func RunPuller() {
	fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	interval := flag.String("interval", "5m", "data puling interval, example: 5m, 30m, 1h")
	help := flag.Bool("h", false, "show usage and exit")
	if err := fset.Parse(os.Args[1:]); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if *help {
		fmt.Fprint(os.Stdout, pullerUsage)
	}

	log := log.New(os.Stdout, "PULLER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	p, err := NewPuller(
		WithInterval(*interval),
	)
	if err != nil {
		panic(err)
	}

	log.Println("puller: Start pulling sensor data")

	for range time.NewTicker(p.Interval).C {
		log.Println("puller: Pulling data")
		if err := p.Run(); err != nil {
			break
		}
		log.Printf("puller: Data saved. Resuming in %s", *interval)
	}
}

var pullerUsage = `
waterlevel - water levels sensor data collector.

Flags:
-h            "Show help"
-interval     "Pulling data interval, example: 1m, 5m, 1h"

Examples:
	// Start puller and collect data every 5 minutes (default settings)
	waterlevel
	// Start puller and collect data every hour
	waterlevel -interval 1h
`
