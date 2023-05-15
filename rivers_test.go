package rivers_test

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"
)

func TestLoadCSV_LoadsExistingFile(t *testing.T) {
	t.Parallel()

	got, err := rivers.LoadWaterLevelCSV("testdata/data.csv")
	if err != nil {
		t.Fatal(err)
	}

	gotRows := len(got)
	wantRows := 96

	if wantRows != gotRows {
		t.Errorf("want %d, got %v", wantRows, gotRows)
	}
}

func TestLoadCSV_FailsOnNotExistingFile(t *testing.T) {
	t.Parallel()
	_, err := rivers.LoadWaterLevelCSV("testdata/notexisting.csv")
	if err == nil {
		t.Error("want error on not existing file")
	}
}

func TestReadCSV(t *testing.T) {
	got, err := rivers.ReadWaterLevelCSV(strings.NewReader(stationData))
	if err != nil {
		t.Fatal(err)
	}

	wantLen := 5
	if wantLen != len(got) {
		t.Errorf("want %d items, got %d", wantLen, len(got))
	}

	wantTimestamp, err := time.Parse(time.RFC3339, "2021-02-10T13:00:00Z")
	if err != nil {
		t.Fatal(err)
	}

	wantValue := 1772

	want := rivers.WaterLevelReading{
		Timestamp: wantTimestamp,
		Value:     wantValue,
	}

	if !cmp.Equal(want, got[0]) {
		t.Errorf(cmp.Diff(want, got[0]))
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
	want := []rivers.WaterLevelReading{
		{
			Name:      "John's Bridge Nore",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     466,
		},
		{
			Name:      "Dinin Bridge",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     53,
		},
		{
			Name:      "Brownsbarn",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     413,
		},
		{
			Name:      "Mount Juliet",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     451,
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
	want := []rivers.WaterLevelReading{
		{
			Name:      "John's Bridge Nore",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     466,
		},
		{
			Name:      "Dinin Bridge",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     53,
		},
		{
			Name:      "Brownsbarn",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     413,
		},
		{
			Name:      "Mount Juliet",
			Timestamp: time.Date(2021, time.June, 15, 22, 00, 00, 00, time.UTC),
			Value:     451,
		},
		{
			Name:      "John's Bridge Nore",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     400,
		},
		{
			Name:      "Dinin Bridge",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     500,
		},
		{
			Name:      "Brownsbarn",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     400,
		},
		{
			Name:      "Mount Juliet",
			Timestamp: time.Date(2021, time.June, 15, 22, 15, 00, 00, time.UTC),
			Value:     400,
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

	stationData = `datetime,value
2021-02-10 13:00,1.772
2021-02-10 13:15,1.771
2021-02-10 13:30,1.769
2021-02-10 13:45,1.769
2021-02-10 14:00,1.768`
)
