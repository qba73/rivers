package rivers_test

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"
)

func TestSaveWritesDataToFile(t *testing.T) {
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

	err = rivers.Save(records, s)
	if err != nil {
		t.Fatal(err)
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

func TestStoreSaveDataToFile(t *testing.T) {
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
}
