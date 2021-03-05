package river_test

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers/internal/river"
	"github.com/qba73/rivers/internal/river/testhelper"
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

func TestLoadCSV(t *testing.T) {

}
