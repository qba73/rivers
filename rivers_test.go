package rivers_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"
)

func TestLoadCSV(t *testing.T) {

	data, err := rivers.LoadCSV("testdata/data.csv")

	if err != nil {
		t.Fatal(err)
	}

	wantLen := 96

	if wantLen != len(data) {
		t.Errorf("want %d items, got %d", wantLen, len(data))
	}

	wantTimestamp, err := time.Parse(time.RFC3339, "2021-02-10T13:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	wantValue := 1.772

	want := rivers.Level{
		Timestamp: wantTimestamp,
		Value:     wantValue,
	}

	if !cmp.Equal(want, data[0]) {
		t.Errorf(cmp.Diff(want, data[0]))
	}

}
