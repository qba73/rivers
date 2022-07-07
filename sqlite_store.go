package rivers

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // DB diver for SQLite3
)

// SQLiteStore represents a data store.
type SQLiteStore struct {
	DB *sql.DB
}

// NewSQLiteStore takes a path and creates a new SQLite store.
// It errors if the filepath is empty.
func NewSQLiteStore(path string) (*SQLiteStore, error) {
	if path == "" {
		return nil, errors.New("empty file path")
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &SQLiteStore{
		DB: db,
	}, nil
}

// Save takes a record representing StationWaterLevelReading and saves it in the store.
func (s *SQLiteStore) Save(record StationWaterLevelReading) error {
	const query = `INSERT INTO waterlevel_readings (station_id, station_name, datetime, value) VALUES (?, ?, datetime(?), ?)`
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, record.StationID, record.Name, record.Readtime, record.WaterLevel)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// List returns all water level reading recorded in the database.
func (s *SQLiteStore) List() ([]StationWaterLevelReading, error) {
	const query = `SELECT station_id, station_name, datetime, value FROM waterlevel_readings`
	rows, err := s.DB.Query(query)
	if err != nil {
		return []StationWaterLevelReading{}, fmt.Errorf("executing DB query: %w", err)
	}
	defer rows.Close()

	var readings []*WaterLevel
	for rows.Next() {
		wl := new(WaterLevel)
		err := rows.Scan(&wl.StationID, &wl.StationName, &wl.Datetime, &wl.Value)
		if err != nil {
			return []StationWaterLevelReading{}, fmt.Errorf("scanning row: %w", err)
		}
		readings = append(readings, wl)
	}
	if err := rows.Err(); err != nil {
		return []StationWaterLevelReading{}, err
	}

	var stationsReadings []StationWaterLevelReading
	for _, r := range readings {
		readTime, err := parseDatetime(r.Datetime)
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
	var wl WaterLevel
	const query = `SELECT station_id, station_name, datetime, value FROM waterlevel_readings WHERE station_id=? order by datetime desc limit 1`

	err := s.DB.QueryRow(query, stationID).Scan(&wl.StationID, &wl.StationName, &wl.Datetime, &wl.Value)
	if errors.Is(err, sql.ErrNoRows) {
		return StationWaterLevelReading{}, fmt.Errorf("no results for stationID %d: %w", stationID, ErrNoReading)
	}
	if err != nil {
		return StationWaterLevelReading{}, fmt.Errorf("selecting last water level reading for stationID %d: %w", stationID, err)
	}

	//	Mon Jan 2 15:04:05 MST 2006
	readTime, err := parseDatetime(wl.Datetime)
	if err != nil {
		return StationWaterLevelReading{}, err
	}
	return StationWaterLevelReading{
		StationID:  wl.StationID,
		Name:       wl.StationName,
		Readtime:   readTime,
		WaterLevel: wl.Value,
	}, nil
}

func parseDatetime(date string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", date)
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
