// +build !integration

package rivers_test

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"
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
	t.Parallel()

	tt := []struct {
		name        string
		filePath    string
		want        int
		expectedErr bool
	}{
		{name: "load correct file", filePath: "testdata/data.csv", want: 96, expectedErr: false},
		{name: "load not existing file", filePath: "testdata/notexisting.csv", want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := rivers.LoadCSV(tc.filePath)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("%s, LoadCSV(%q) failed, error: %v", tc.name, tc.filePath, err)
			}

			if !tc.expectedErr && (tc.want != len(got)) {
				t.Errorf("%s, LoadCSV(%q) got: %v, want: %d", tc.name, tc.filePath, got, tc.want)
			}
		})
	}
}

func TestReadCSV(t *testing.T) {
	file := setup(t)
	defer cleanup(file)

	csvFile, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("can't read csv data: %s", err)
	}

	data, err := rivers.ReadCSV(csvFile)
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

	want := rivers.SensorReading{
		Timestamp: wantTimestamp,
		Value:     wantValue,
	}

	if !cmp.Equal(want, data[0]) {
		t.Errorf(cmp.Diff(want, data[0]))
	}
}
