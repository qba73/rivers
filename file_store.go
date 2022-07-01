package rivers

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

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
	f, err := os.OpenFile(fs.path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o600)
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
