package rivers

// Store is the interface that wraps Save,
// GetLastReadingForStationID and List methods.
type Store interface {
	// Save takes a record and stores it in the store.
	Save(StationWaterLevelReading) error

	// GetLastReadingsForStationID takes station ID and returns
	// last recorded reading.
	GetLastReadingForStationID(int) (StationWaterLevelReading, error)

	// List returns all readings from the store.
	List() ([]StationWaterLevelReading, error)
}
