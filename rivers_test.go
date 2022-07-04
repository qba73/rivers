//go:build !integration
// +build !integration

package rivers_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"
	"github.com/qba73/rivers/testhelper"
)

func setup(t *testing.T) *os.File {
	file := testhelper.TmpTextFile(t, "", "data.csv", `datetime,value
2021-02-10 13:00,1.772
2021-02-10 13:15,1.771
2021-02-10 13:30,1.769
2021-02-10 13:45,1.769
2021-02-10 14:00,1.768`)

	t.Cleanup(func() {
		os.Remove(file.Name())
	})
	return file
}

func TestLoadCSV(t *testing.T) {
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
			t.Parallel()
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

	want := rivers.Reading{
		Timestamp: wantTimestamp,
		Value:     wantValue,
	}

	if !cmp.Equal(want, data[0]) {
		t.Errorf(cmp.Diff(want, data[0]))
	}
}

func TestParseStationGroup_ErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := rivers.ReadGroupCSV(strings.NewReader(invalidGroupInputNoData))
	if err == nil {
		t.Fatal("want err on empty input")
	}
}

func TestParseStationGroup_ErrorsOnHeaderOnlyCSV(t *testing.T) {
	t.Parallel()
	_, err := rivers.ReadGroupCSV(strings.NewReader(invalidGroupInputHeaderOnly))
	if err == nil {
		t.Fatal("want err on empty input")
	}
}

func TestParseStationGroup_ErrorsOnNoStationInCSV(t *testing.T) {
	t.Parallel()
	_, err := rivers.ReadGroupCSV(strings.NewReader(invalidGroupInputOneColumnOnly))
	if err == nil {
		t.Fatal("want err on a record without station")
	}
}

func TestParseStationGroup_ParsesSingleRecord(t *testing.T) {
	t.Parallel()
	want := []rivers.Reading{
		{
			Name:      "John's Bridge Nore",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.466,
		},
		{
			Name:      "Dinin Bridge",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.053,
		},
		{
			Name:      "Brownsbarn",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.413,
		},
		{
			Name:      "Mount Juliet",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.451,
		},
	}

	got, err := rivers.ReadGroupCSV(strings.NewReader(validGroupInputSingleRecord))
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestParseStationGroup_ParsesMultipleRecords(t *testing.T) {
	t.Parallel()
	want := []rivers.Reading{
		{
			Name:      "John's Bridge Nore",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.466,
		},
		{
			Name:      "Dinin Bridge",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.053,
		},
		{
			Name:      "Brownsbarn",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.413,
		},
		{
			Name:      "Mount Juliet",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     0.451,
		},
		{
			Name:      "John's Bridge Nore",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     0.400,
		},
		{
			Name:      "Dinin Bridge",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     0.500,
		},
		{
			Name:      "Brownsbarn",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     0.400,
		},
		{
			Name:      "Mount Juliet",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     0.400,
		},
	}

	got, err := rivers.ReadGroupCSV(strings.NewReader(validGroupInputMultipleRecords))
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

var (
	validGroupInputSingleRecord = `Datetime,John's Bridge Nore,Dinin Bridge,Brownsbarn,Mount Juliet
2021-06-15 22:00,0.466,0.053,0.413,0.451`
	validGroupInputMultipleRecords = `Datetime,John's Bridge Nore,Dinin Bridge,Brownsbarn,Mount Juliet
2021-06-15 22:00,0.466,0.053,0.413,0.451
2021-06-15 22:15,0.400,0.500,0.400,0.400`
	invalidGroupInputNoData        = ``
	invalidGroupInputHeaderOnly    = `Datetime,John's Bridge Nore,Dinin Bridge,Brownsbarn,Mount Juliet`
	invalidGroupInputOneColumnOnly = `Datetime
2021-06-15 22:00`
)
