package river_test

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	river "github.com/qba73/rivers"
	"github.com/qba73/rivers/testhelper"
)

func setup(t *testing.T) *os.File {
	return testhelper.TmpTextFile(t, "", "data.csv", `datetime,value
2021-02-10 13:00,1.772
2021-02-10 13:15,1.771
2021-02-10 13:30,1.769
2021-02-10 13:45,1.769
2021-02-10 14:00,1.768`)
}

func cleanup(file *os.File) {
	os.Remove(file.Name())
}

func TestLoadCSV(t *testing.T) {
	t.Run("Load correct file", func(t *testing.T) {
		testFile := "testdata/data.csv"
		got, err := river.LoadCSV(testFile)
		if err != nil {
			t.Fatalf("LoadCSV(%s) returned error: %s", testFile, err)
		}

		wantLen := 96
		if wantLen != len(got) {
			t.Errorf("got: %d, want: %d", len(got), wantLen)
		}
	})

	t.Run("Load not existing file", func(t *testing.T) {
		testFile := "testdata/notexisting.csv"
		expectedErr := true

		got, err := river.LoadCSV(testFile)
		if (err != nil) != expectedErr {
			t.Fatalf("LoadCSV(%s) returned error: %s", testFile, err)
		}

		if !expectedErr && (got != nil) {
			t.Errorf("got %v, want nil", got)
		}
	})
}

func TestReadCSV(t *testing.T) {
	file := setup(t)
	defer cleanup(file)

	csvFile, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("can't read csv data: %s", err)
	}

	data, err := river.ReadCSV(csvFile)
	if err != nil {
		t.Fatal(err)
	}

	wantLen := 5
	if wantLen != len(data) {
		t.Errorf("want %d items, got %d", wantLen, len(data))
	}

	wantTimestamp, err := time.Parse(time.RFC3339, "2021-02-10T13:00:00Z")
	if err != nil {
		t.Fatal(err)
	}

	wantValue := 1.772

	want := river.Level{
		Timestamp: wantTimestamp,
		Value:     wantValue,
	}

	if !cmp.Equal(want, data[0]) {
		t.Errorf(cmp.Diff(want, data[0]))
	}
}
