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

func WithLogger(l *log.Logger) option {
	return func(p *Puller) error {
		p.Log = l
		return nil
	}
}

type Puller struct {
	Client      *Client
	ReadingRepo *ReadingsRepo
	Interval    time.Duration
	Log         *log.Logger
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

func (p *Puller) RunPeriodically() error {
	p.Log.Printf("puller : Start water levels puller with interval: %s", p.Interval)
	for range time.NewTicker(p.Interval).C {
		p.Log.Println("puller : Pull latest water levels")
		stationReadings, err := p.Client.GetLatestWaterLevels()
		if err != nil {
			return fmt.Errorf("retriving water level data: %w", err)
		}
		for _, reading := range stationReadings {
			err = p.ReadingRepo.Add(reading)
			if errors.Is(err, ErrReadingExists) {
				continue
			}
		}
		p.Log.Printf("puller : Saved latest water levels, resuming in %s", p.Interval.String())
	}
	return nil
}

// RunPuller holds all required machinery to run the water levels data puller.
func RunPuller() {
	fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	interval := fset.String("interval", "5m", "data pulling interval, example: 5m, 30m, 1h")
	help := fset.Bool("h", false, "show usage and exit")
	if err := fset.Parse(os.Args[1:]); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if *help {
		fmt.Fprint(os.Stdout, pullerUsage)
		os.Exit(0)
	}

	log := log.New(os.Stdout, "PULLER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	p, err := NewPuller(
		WithInterval(*interval),
		WithLogger(log),
	)
	if err != nil {
		panic(err)
	}

	if err := p.RunPeriodically(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var pullerUsage string = `
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
