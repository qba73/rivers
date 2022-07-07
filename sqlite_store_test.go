package rivers_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQLStore_RetrievesLastReadingForGivenStation(t *testing.T) {
	t.Parallel()
	db := newTestDB(stmtRetrieveLastReadingForOneStation, t)
	store := rivers.SQLiteStore{DB: db}
	got, err := store.GetLastReadingForStationID(1043)
	if err != nil {
		t.Fatal(err)
	}
	want := rivers.StationWaterLevelReading{
		StationID:  1043,
		Name:       "Ballybofey",
		Readtime:   time.Date(2022, 06, 30, 04, 15, 00, 00, time.UTC),
		WaterLevel: 879,
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSQLStore_RetrievesAllReadings(t *testing.T) {
	t.Parallel()
	db := newTestDB(stmtRetrieveLastReadingForOneStation, t)
	store := rivers.SQLiteStore{DB: db}
	got, err := store.List()
	if err != nil {
		t.Fatal(err)
	}
	// Number of records in the newly provisioned test DB.
	want := 5
	if want != len(got) {
		t.Errorf("want %d records, got %d", want, len(got))
	}
}
