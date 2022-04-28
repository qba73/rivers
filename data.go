package rivers

import (
	"encoding/json"
	"errors"
	"os"
)

// Saver is an interface that wraps the Save method.
type Saver interface {
	Save(records []StationWaterLevelReading) error
}

// FileStore is a data store.
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

// Save takes a slice of records and saves them in the file.
func (fs FileStore) Save(records []StationWaterLevelReading) error {
	f, err := os.Create(fs.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(records)
}

// Save knows how to save records using provided saver.
func Save(records []StationWaterLevelReading, saver Saver) error {
	return saver.Save(records)
}
