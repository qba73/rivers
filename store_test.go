package rivers_test

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
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
	db, err := rivers.Open("testdata/water.db")
	if err != nil {
		t.Fatal(err)
	}
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
			Value:       0.384,
		},
		{
			StationID:   1043,
			StationName: "Ballybofey",
			SensorRef:   "0001",
			Datetime:    "2022-06-28T04:14:00Z",
			Value:       1.679,
		},
		{
			StationID:   3055,
			StationName: "Glaslough",
			Datetime:    "2022-06-28T04:45:00Z",
			SensorRef:   "0001",
			Value:       0.4779999999999999,
		},
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestRetrieveLastReadingForOneStation(t *testing.T) {
	t.Parallel()
	db, err := rivers.Open("testdata/water.db")
	if err != nil {
		t.Fatal(err)
	}
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
		Datetime:    "2022-06-28T04:14:00Z",
		Value:       1.679,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
