package rivers_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/qba73/rivers"
)

func TestListGetsAllWaterLevelReadingsFromDatabase(t *testing.T) {
	t.Parallel()
	db := createTestDB(waterLevelSchema, t)
	prepareTestData(stmtRetrieveLastReadingForOneStation, db, t)
	defer cleanupTestDB(dropWaterLevelSchema, db, t)

	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}
	got, err := readings.List()
	if err != nil {
		t.Fatal(err)
	}
	want := []rivers.WaterLevel{
		{
			StationID:   1042,
			StationName: "Sandy Millss",
			Datetime:    "2022-06-28 04:45:00-00:00",
			Value:       0.383,
		},
		{
			StationID:   1043,
			StationName: "Ballybofey",
			Datetime:    "2022-06-28 04:15:00-00:00",
			Value:       1.679,
		},
		{
			StationID:   1043,
			StationName: "Ballybofey",
			Datetime:    "2022-06-29 05:15:00-00:00",
			Value:       1.779,
		},
		{
			StationID:   1043,
			StationName: "Ballybofey",
			Datetime:    "2022-06-30 04:15:00-00:00",
			Value:       1.879,
		},
		{
			StationID:   3055,
			StationName: "Glaslough",
			Datetime:    "2022-06-28 04:45:00-00:00",
			Value:       0.478,
		},
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestRetrieveLastReadingForOneStation(t *testing.T) {
	t.Parallel()
	db := createTestDB(waterLevelSchema, t)
	prepareTestData(stmtRetrieveLastReadingForOneStation, db, t)
	defer cleanupTestDB(dropWaterLevelSchema, db, t)

	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}
	got, err := readings.GetLastReadingForStationID("1043")
	if err != nil {
		t.Fatal(err)
	}
	readTime, err := time.Parse("2006-01-02 15:04:05-07:00", "2022-06-30 04:15:00-00:00")
	if err != nil {
		t.Fatal(err)
	}
	want := rivers.StationWaterLevelReading{
		StationID:  "1043",
		Name:       "Ballybofey",
		Readtime:   readTime,
		WaterLevel: 1.879,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddSingleReadingToTheStore(t *testing.T) {
	t.Parallel()
	db := createTestDB(waterLevelSchema, t)
	defer cleanupTestDB(dropWaterLevelSchema, db, t)

	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}
	want := rivers.StationWaterLevelReading{
		StationID:  "3055",
		Name:       "Glaslough",
		Readtime:   time.Now(),
		WaterLevel: 0.991,
	}

	err := readings.Add(want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := readings.GetLastReadingForStationID("3055")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddLatest_DoesNotAddDuplicateReadings(t *testing.T) {
	t.Parallel()
	db := createTestDB(waterLevelSchema, t)
	defer cleanupTestDB(dropWaterLevelSchema, db, t)
	readings := rivers.ReadingsRepo{
		Store: &rivers.SQLiteStore{
			DB: db,
		},
	}

	want := rivers.StationWaterLevelReading{
		StationID: "3055",
		Readtime:  time.Now(),
	}
	err := readings.Add(want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := readings.GetLastReadingForStationID("3055")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	// Try to add the same reading second time.
	err = readings.AddLatest(want)
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

func createTestDB(stmt string, t *testing.T) *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	_, err := db.Exec(stmt)
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func prepareTestData(stmt string, db *sqlx.DB, t *testing.T) {
	_, err := db.Exec(stmt)
	if err != nil {
		t.Fatal(err)
	}
}

func cleanupTestDB(stmt string, db *sqlx.DB, t *testing.T) {
	_, err := db.Exec(stmt)
	if err != nil {
		t.Fatal(err)
	}
}

var (
	waterLevelSchema = `DROP TABLE IF EXISTS waterlevel_readings;
CREATE TABLE waterlevel_readings (
id INTEGER PRIMARY KEY AUTOINCREMENT,
station_id INT NOT NULL,
station_name CHAR(50) NOT NULL,
datetime TEXT NOT NULL,
value REAL);`

	dropWaterLevelSchema = `DROP TABLE waterlevel_readings`

	// DB statements for populating data
	stmtRetrieveLastReadingForOneStation = `INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES (1042,'Sandy Millss','2022-06-28 04:45:00-00:00',0.383);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(1043,'Ballybofey','2022-06-28 04:15:00-00:00',1.679);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(1043,'Ballybofey','2022-06-29 05:15:00-00:00',1.779);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(1043,'Ballybofey','2022-06-30 04:15:00-00:00',1.879);
INSERT INTO "waterlevel_readings" (station_id, station_name, datetime, value) VALUES(3055,'Glaslough','2022-06-28 04:45:00-00:00',0.478);`
)
