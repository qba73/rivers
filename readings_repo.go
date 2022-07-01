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

// List returns all water level readings.
func (r *ReadingsRepo) List() ([]WaterLevel, error) {
	var readings []WaterLevel
	const query = `SELECT station_id, station_name, datetime, value FROM waterlevel_readings`
	if err := r.Store.(*SQLiteStore).DB.Select(&readings, query); err != nil {
		return nil, fmt.Errorf("selecting water level readings: %w", err)
	}
	return readings, nil
}

// GetLastReadingForStationID retrieves latest water level reading for given station id.
func (r *ReadingsRepo) GetLastReadingForStationID(strStationID string) (StationWaterLevelReading, error) {
	// could do more here!
	return r.Store.GetLastReadingForStationID(strStationID)
}

// Add takes a reading and adds it to the store.
func (r *ReadingsRepo) Add(level StationWaterLevelReading) error {
	// business logic here...?
	return r.Store.Save(level)
}

// AddLatest takes water lever readings and add it to the store
// if the reading does not already exist. It returns an error
// if the reading already is present in the store.
func (r *ReadingsRepo) AddLatest(level StationWaterLevelReading) error {
	current, err := r.GetLastReadingForStationID(level.StationID)
	if err != nil && !errors.Is(err, ErrNoReading) {
		return err
	}
	if errors.Is(err, ErrNoReading) {
		return r.Add(level)
	}
	if current.StationID == level.StationID && current.Readtime.Equal(level.Readtime) {
		return fmt.Errorf("adding sensor reading %v: %w", level, ErrReadingExists)
	}
	return r.Add(level)
}

var (
	ErrReadingExists = errors.New("duplicated sensor entry")
	ErrNoReading     = errors.New("reading entry doesn't exist")
)
