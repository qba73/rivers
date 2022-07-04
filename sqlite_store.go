package rivers

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// SQLiteStore represents a data store.
type SQLiteStore struct {
	DB *sqlx.DB
}

// NewSQLiteStore takes a path and creates a new SQLite store.
// It errors if the filepath is empty.
func NewSQLiteStore(path string) (*SQLiteStore, error) {
	if path == "" {
		return nil, errors.New("empty file path")
	}
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &SQLiteStore{
		DB: db,
	}, nil
}

// Save takes a record representing StationWaterLevelReading and saves it in the store.
func (s *SQLiteStore) Save(record StationWaterLevelReading) error {
	levelreadings := `INSERT INTO waterlevel_readings (station_id, station_name, datetime, value) VALUES (?, ?, ?, ?)`
	_, err := s.DB.Exec(levelreadings, record.StationID, record.Name, record.Readtime, record.WaterLevel)
	if err != nil {
		return fmt.Errorf("adding waterlevel reading %#v: %w", record, err)
	}
	return nil
}

func (s *SQLiteStore) List() ([]StationWaterLevelReading, error) {
	var readings []WaterLevel
	const query = `SELECT station_id, station_name, datetime, value FROM waterlevel_readings`
	if err := s.DB.Select(&readings, query); err != nil {
		return nil, fmt.Errorf("selecting water level readings: %w", err)
	}
	var stationsReadings []StationWaterLevelReading
	for _, r := range readings {
		readTime, err := time.Parse("2006-01-02 15:04:05-07:00", r.Datetime)
		if err != nil {
			return []StationWaterLevelReading{}, err
		}
		reading := StationWaterLevelReading{
			StationID:  r.StationID,
			Name:       r.StationName,
			Readtime:   readTime,
			WaterLevel: r.Value,
		}
		stationsReadings = append(stationsReadings, reading)
	}
	return stationsReadings, nil
}

// GetLastReadingForStationID retrieves latest water level reading for given station id.
func (s *SQLiteStore) GetLastReadingForStationID(stationID int) (StationWaterLevelReading, error) {
	var reading []WaterLevel
	levelreading := `SELECT station_id, station_name, datetime, value FROM waterlevel_readings WHERE station_id=? order by datetime desc limit 1`
	if err := s.DB.Select(&reading, levelreading, stationID); err != nil {
		return StationWaterLevelReading{}, fmt.Errorf("selecting last water level reading for stationID %d: %w", stationID, err)
	}
	if len(reading) == 0 {
		return StationWaterLevelReading{}, fmt.Errorf("no results for stationID %d: %w", stationID, ErrNoReading)
	}
	r := reading[0]
	//	Mon Jan 2 15:04:05 MST 2006
	readTime, err := time.Parse("2006-01-02 15:04:05-07:00", r.Datetime)
	if err != nil {
		return StationWaterLevelReading{}, err
	}
	return StationWaterLevelReading{
		StationID:  r.StationID,
		Name:       r.StationName,
		Readtime:   readTime,
		WaterLevel: r.Value,
	}, nil
}

// WaterLevel represents water level value
// recorded at the given time by the sensor installed
// in the station.
type WaterLevel struct {
	StationID   int    `db:"station_id"`
	StationName string `db:"station_name"`
	Datetime    string `db:"datetime"`
	Value       int    `db:"value"`
}
