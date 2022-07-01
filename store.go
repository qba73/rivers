package rivers

import (
	_ "github.com/mattn/go-sqlite3"
)

// Store is the interface that wraps the basic Save method.
//
// Save takes records and stores them in a store.
type Store interface {
	Save(StationWaterLevelReading) error
	GetLastReadingForStationID(string) (StationWaterLevelReading, error)
}
