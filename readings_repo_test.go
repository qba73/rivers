package rivers_test

import (
	"errors"
	"testing"
	"time"

	"database/sql"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	_ "github.com/mattn/go-sqlite3" // DB diver for SQLite3
	"github.com/qba73/rivers"
)

func TestListGetsAllWaterLevelReadingsFromDatabase(t *testing.T) {
	t.Parallel()
	db := newTestDB(stmtRetrieveLastReadingForOneStation, t)

	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}
	got, err := readings.List()
	if err != nil {
		t.Fatal(err)
	}
	want := []rivers.StationWaterLevelReading{
		{
			StationID:  1042,
			Name:       "Sandy Millss",
			Readtime:   time.Date(2022, 06, 28, 04, 45, 00, 00, time.UTC),
			WaterLevel: 383,
		},
		{
			StationID:  1043,
			Name:       "Ballybofey",
			Readtime:   time.Date(2022, 06, 28, 04, 15, 00, 00, time.UTC),
			WaterLevel: 679,
		},
		{
			StationID:  1043,
			Name:       "Ballybofey",
			Readtime:   time.Date(2022, 06, 29, 05, 15, 00, 00, time.UTC),
			WaterLevel: 779,
		},
		{
			StationID:  1043,
			Name:       "Ballybofey",
			Readtime:   time.Date(2022, 06, 30, 04, 15, 00, 00, time.UTC),
			WaterLevel: 879,
		},
		{
			StationID:  3055,
			Name:       "Glaslough",
			Readtime:   time.Date(2022, 06, 28, 04, 45, 00, 00, time.UTC),
			WaterLevel: 478,
		},
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestRetrieveLastReadingForOneStation(t *testing.T) {
	t.Parallel()
	db := newTestDB(stmtRetrieveLastReadingForOneStation, t)
	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}
	got, err := readings.GetLastReadingForStationID(1043)
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

func TestAddSingleReadingToTheStore(t *testing.T) {
	t.Parallel()
	db := newTestDB(stmtEmptyDB, t)
	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}
	want := rivers.StationWaterLevelReading{
		StationID:  3055,
		Name:       "Glaslough",
		Readtime:   time.Now(),
		WaterLevel: 991,
	}
	err := readings.Add(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := readings.GetLastReadingForStationID(3055)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreFields(rivers.StationWaterLevelReading{}, "Readtime")) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAdd_DoesNotAddDuplicateReadings(t *testing.T) {
	t.Parallel()
	db := newTestDB(stmtEmptyDB, t)
	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}
	want := rivers.StationWaterLevelReading{
		StationID: 3055,
		Readtime:  time.Date(2022, 06, 28, 04, 15, 00, 00, time.UTC),
	}
	err := readings.Add(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := readings.GetLastReadingForStationID(3055)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreFields(rivers.StationWaterLevelReading{}, "Readtime")) {
		t.Error(cmp.Diff(want, got))
	}
	// Try to add the same reading second time.
	err = readings.Add(want)
	if !errors.Is(err, rivers.ErrReadingExists) {
		t.Fatalf("want ErrReadingExists, got %v", err)
	}
	gotReadings, err := readings.List()
	if err != nil {
		t.Error(err)
	}
	if len(gotReadings) != 1 {
		t.Error("it's not filtering out duplicates")
	}
}

func newTestDB(stmtPopulateData string, t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(waterLevelSchema)
	if err != nil {
		t.Fatal(err)
	}
	if stmtPopulateData != "" {
		_, err = db.Exec(stmtPopulateData)
		if err != nil {
			t.Fatal(err)
		}
	}
	// Call the func to clean database after each test.
	// This way we don't need to pullute test logic with `defer`.
	t.Cleanup(func() {
		if _, err := db.Exec(`DROP TABLE waterlevel_readings`); err != nil {
			t.Fatalf("error cleaning up test database: %#v", err)
		}
	})
	return db
}

var (
	waterLevelSchema = `DROP TABLE IF EXISTS waterlevel_readings;
CREATE TABLE waterlevel_readings (
id INTEGER PRIMARY KEY AUTOINCREMENT,
station_id INT NOT NULL,
station_name CHAR(50) NOT NULL,
datetime TEXT NOT NULL,
value INTEGER);`

	// DB statements for populating data
	stmtRetrieveLastReadingForOneStation = `INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES (1042,'Sandy Millss',datetime('2022-06-28 04:45:00-00:00'),383);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(1043,'Ballybofey',datetime('2022-06-28 04:15:00-00:00'),679);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(1043,'Ballybofey',datetime('2022-06-29 05:15:00-00:00'),779);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(1043,'Ballybofey',datetime('2022-06-30 04:15:00-00:00'),879);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(3055,'Glaslough',datetime('2022-06-28 04:45:00-00:00'),478);`
	stmtEmptyDB = ``
)
