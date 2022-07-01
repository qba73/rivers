package rivers_test

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/qba73/rivers"
)

func TestFileStore_SavesDataToFile(t *testing.T) {
	t.Parallel()
	records := []rivers.StationWaterLevelReading{
		{
			StationID:  "0000001041",
			Readtime:   time.Date(2021, 02, 18, 06, 00, 00, 00, time.UTC),
			WaterLevel: 1.715,
		},
		{
			StationID:  "0000001042",
			Readtime:   time.Date(2021, 02, 18, 07, 00, 00, 00, time.UTC),
			WaterLevel: 2.715,
		},
	}
	path := t.TempDir() + "/data_test.txt"

	s, err := rivers.NewFileStore(path)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Save(records)
	if err != nil {
		t.Fatalf("save returned error: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want, err := os.ReadFile("testdata/savedata_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFileStore_AppendsDataToFile(t *testing.T) {
	t.Parallel()
	records := []rivers.StationWaterLevelReading{
		{
			StationID:  "0000001041",
			Readtime:   time.Date(2021, 02, 18, 06, 00, 00, 00, time.UTC),
			WaterLevel: 1.715,
		},
		{
			StationID:  "0000001042",
			Readtime:   time.Date(2021, 02, 18, 07, 00, 00, 00, time.UTC),
			WaterLevel: 2.715,
		},
	}

	path := t.TempDir() + "/data_test.txt"

	s, err := rivers.NewFileStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Save(records)

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want, err := os.ReadFile("testdata/savedata_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	// Append new reading records to the file
	newRecords := []rivers.StationWaterLevelReading{
		{
			StationID:  "0000001051",
			Readtime:   time.Date(2021, 02, 18, 06, 00, 00, 00, time.UTC),
			WaterLevel: 2.715,
		},
		{
			StationID:  "0000001052",
			Readtime:   time.Date(2021, 02, 18, 07, 00, 00, 00, time.UTC),
			WaterLevel: 3.715,
		},
	}
	s.Save(newRecords)
	got, err = os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want, err = os.ReadFile("testdata/appenddata_test.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestLoadData_ReadsAllRecordsFromFile(t *testing.T) {
	t.Parallel()

	want := []rivers.StationWaterLevelReading{
		{
			StationID:  "0000001041",
			Readtime:   time.Date(2021, 02, 18, 06, 00, 00, 00, time.UTC),
			WaterLevel: 1.715,
		},
		{
			StationID:  "0000001042",
			Readtime:   time.Date(2021, 02, 18, 07, 00, 00, 00, time.UTC),
			WaterLevel: 2.715,
		},
		{
			StationID:  "0000001051",
			Readtime:   time.Date(2021, 02, 18, 06, 00, 00, 00, time.UTC),
			WaterLevel: 2.715,
		},
		{
			StationID:  "0000001052",
			Readtime:   time.Date(2021, 02, 18, 07, 00, 00, 00, time.UTC),
			WaterLevel: 3.715,
		},
	}

	path := "testdata/appenddata_test.txt"
	s, err := rivers.NewFileStore(path)
	if err != nil {
		t.Fatal(err)
	}

	got, err := s.Records()
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestListGetsAllWaterLevelReadingsFromDatabase(t *testing.T) {
	t.Parallel()
	db := createTestDB(waterLevelSchema, t)
	prepareTestData(stmtRetrieveLastReadingForOneStation, db, t)
	defer cleanupTestDB(dropWaterLevelSchema, db, t)

	levels := rivers.Readings{DB: db}
	got, err := levels.List()
	if err != nil {
		t.Fatal(err)
	}
	want := []rivers.WaterLevel{
		{
			StationID:   1042,
			StationName: "Sandy Millss",
			SensorRef:   "0001",
			Datetime:    "2022-06-28T04:45:00Z",
			Value:       0.383,
		},
		{
			StationID:   1043,
			StationName: "Ballybofey",
			SensorRef:   "0001",
			Datetime:    "2022-06-28T04:15:00Z",
			Value:       1.679,
		},
		{
			StationID:   1043,
			StationName: "Ballybofey",
			SensorRef:   "0001",
			Datetime:    "2022-06-29T05:15:00Z",
			Value:       1.779,
		},
		{
			StationID:   1043,
			StationName: "Ballybofey",
			SensorRef:   "0001",
			Datetime:    "2022-06-30T04:15:00Z",
			Value:       1.879,
		},
		{
			StationID:   3055,
			StationName: "Glaslough",
			SensorRef:   "0001",
			Datetime:    "2022-06-28T04:45:00Z",
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

	readings := rivers.Readings{DB: db}
	stationID := 1043
	got, err := readings.GetLast(stationID)
	if err != nil {
		t.Fatal(err)
	}
	want := rivers.WaterLevel{
		StationID:   1043,
		StationName: "Ballybofey",
		SensorRef:   "0001",
		Datetime:    "2022-06-30T04:15:00Z",
		Value:       1.879,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddSingleReadingToTheStore(t *testing.T) {
	t.Parallel()
	db := createTestDB(waterLevelSchema, t)
	defer cleanupTestDB(dropWaterLevelSchema, db, t)

	readings := rivers.Readings{DB: db}

	want := rivers.WaterLevel{
		StationID:   3055,
		StationName: "Glaslough",
		SensorRef:   "0001",
		Datetime:    "2022-06-30T19:15:00Z",
		Value:       0.991,
	}

	err := readings.Add(want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := readings.GetLast(3055)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

var (
	waterLevelSchema = `DROP TABLE IF EXISTS waterlevel_readings;
CREATE TABLE waterlevel_readings (
id INTEGER PRIMARY KEY AUTOINCREMENT,
station_id INT NOT NULL,
station_name CHAR(50) NOT NULL,
sensor_ref CHAR(20) NOT NULL,
datetime TEXT NOT NULL,
value REAL);`

	dropWaterLevelSchema = `DROP TABLE waterlevel_readings`

	// DB statements for populating data
	stmtRetrieveLastReadingForOneStation = `INSERT INTO "waterlevel_readings" (station_id, station_name, sensor_ref, datetime, value) VALUES (1042,'Sandy Millss','0001','2022-06-28T04:45:00Z',0.383);
INSERT INTO "waterlevel_readings" (station_id, station_name, sensor_ref, datetime, value) VALUES(1043,'Ballybofey','0001','2022-06-28T04:15:00Z',1.679);
INSERT INTO "waterlevel_readings" (station_id, station_name, sensor_ref, datetime, value) VALUES(1043,'Ballybofey','0001','2022-06-29T05:15:00Z',1.779);
INSERT INTO "waterlevel_readings" (station_id, station_name, sensor_ref, datetime, value) VALUES(1043,'Ballybofey','0001','2022-06-30T04:15:00Z',1.879);
INSERT INTO "waterlevel_readings" (station_id, station_name, sensor_ref, datetime, value) VALUES(3055,'Glaslough','0001','2022-06-28T04:45:00Z',0.478);`
)

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
