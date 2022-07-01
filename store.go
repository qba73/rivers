package rivers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Saver is the interface that wraps the basic Save method.
//
// Save takes records and stores them in a store.
type Saver interface {
	Save(records []StationWaterLevelReading) error
}

// FileStore represents a data store.
type FileStore struct {
	path string
}

// NewFileStore takes a path and creates a new file store.
// It errors if the filepath is empty.
func NewFileStore(path string) (*FileStore, error) {
	if path == "" {
		return nil, errors.New("empty file path")
	}
	fstore := FileStore{
		path: path,
	}
	return &fstore, nil
}

// Save takes a slice of records and saves them in a file.
func (fs *FileStore) Save(records []StationWaterLevelReading) error {
	f, err := os.OpenFile(fs.path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(records)
}

// Records returns records stored in a file.
func (fs *FileStore) Records() ([]StationWaterLevelReading, error) {
	f, err := os.Open(fs.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var records []StationWaterLevelReading
	decoder := json.NewDecoder(f)
	for {
		rec := []StationWaterLevelReading{}
		err := decoder.Decode(&rec)
		if errors.Is(err, io.EOF) {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, rec...)
	}
}

// Open returns database connection.
func Open(path string) (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", path)
}

// WaterLevel represents water level value
// recorded at the given time by the sensor installed
// in the station.
type WaterLevel struct {
	StationID   int     `db:"station_id"`
	StationName string  `db:"station_name"`
	SensorRef   string  `db:"sensor_ref"`
	Datetime    string  `db:"datetime"`
	Value       float64 `db:"value"`
}

type Readings struct {
	DB *sqlx.DB
}

// List returns all water level readings.
func (r *Readings) List() ([]WaterLevel, error) {
	var readings []WaterLevel
	const query = `SELECT station_id, station_name, sensor_ref, datetime, value FROM waterlevel_readings`
	if err := r.DB.Select(&readings, query); err != nil {
		return nil, fmt.Errorf("selecting water level readings: %w", err)
	}
	return readings, nil
}

// GetLast retrieves latest water level reading for given station id.
func (r *Readings) GetLast(stationID int) (WaterLevel, error) {
	var reading []WaterLevel
	levelreading := `SELECT station_id, station_name, sensor_ref, datetime, value FROM waterlevel_readings WHERE station_id=? order by datetime desc limit 1`
	if err := r.DB.Select(&reading, levelreading, stationID); err != nil {
		return WaterLevel{}, fmt.Errorf("selecting last water level reading for stationID %d: %w", stationID, err)
	}
	if len(reading) == 0 {
		return WaterLevel{}, fmt.Errorf("no results for stationID %d: %w", stationID, ErrNoReading)
	}
	return reading[0], nil
}

// Add takes a reading and adds it to the store.
func (r *Readings) Add(level WaterLevel) error {
	levelreadings := `INSERT INTO waterlevel_readings (station_id, station_name, sensor_ref, datetime, value) VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(levelreadings, level.StationID, level.StationName, level.SensorRef, level.Datetime, level.Value)
	if err != nil {
		return fmt.Errorf("adding waterlevel reading %v: %w", level, err)
	}
	return nil
}

// AddLatest takes water lever readings and add it to the store
// if the reading does not already exist. It returns an error
// if the reading already is present in the store.
func (r *Readings) AddLatest(level WaterLevel) error {
	current, err := r.GetLast(level.StationID)
	if err != nil && !errors.Is(err, ErrNoReading) {
		return err
	}
	if errors.Is(err, ErrNoReading) {
		return r.Add(level)
	}
	if current.StationID == level.StationID && current.Datetime == level.Datetime {
		return fmt.Errorf("adding sensor reading %v: %w", level, ErrReadingExists)
	}
	return r.Add(level)
}

var (
	ErrReadingExists = errors.New("duplicated sensor entry")
	ErrNoReading     = errors.New("reading entry doesn't exist")
)
