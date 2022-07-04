package rivers

import (
	"errors"
	"fmt"
)

type ReadingsRepo struct {
	Store Store
}

// OpenReadingsRepo returns database connection.
func OpenReadingsRepo(s Store) *ReadingsRepo {
	return &ReadingsRepo{
		Store: s,
	}
}

// List returns all water level readings from the store.
func (r *ReadingsRepo) List() ([]StationWaterLevelReading, error) {
	return r.Store.List()
}

// GetLastReadingForStationID retrieves latest water level reading for given station id.
func (r *ReadingsRepo) GetLastReadingForStationID(stationID int) (StationWaterLevelReading, error) {
	// could do more here!
	return r.Store.GetLastReadingForStationID(stationID)
}

// Add takes a reading and adds it to the store.
func (r *ReadingsRepo) Add(reading StationWaterLevelReading) error {
	current, err := r.Store.GetLastReadingForStationID(reading.StationID)
	if err != nil && !errors.Is(err, ErrNoReading) {
		return fmt.Errorf("adding sensor reading: %w", err)
	}
	if errors.Is(err, ErrNoReading) {
		return r.Store.Save(reading)
	}
	if current.StationID == reading.StationID && current.Readtime.Equal(reading.Readtime) {
		return fmt.Errorf("adding sensor reading %v: %w", reading, ErrReadingExists)
	}
	return r.Store.Save(current)
}

var (
	ErrReadingExists = errors.New("duplicated sensor entry")
	ErrNoReading     = errors.New("reading entry doesn't exist")
)
