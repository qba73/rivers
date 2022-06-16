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
