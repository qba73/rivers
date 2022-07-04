package rivers

import (
	_ "github.com/mattn/go-sqlite3"
)

// Store is the interface that wraps Save
// and GetLastReadingForStationID methods.
//
type Store interface {
	// Save takes records and stores them in a store.
	Save(StationWaterLevelReading) error

	// GetLastReadingsForStationID takes station ID and returns
	// last recorded reading.
	GetLastReadingForStationID(int) (StationWaterLevelReading, error)
}
