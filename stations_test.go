package rivers_test

import (
	"testing"

	"github.com/qba73/rivers"
)

func TestLoadStations(t *testing.T) {

	data, err := rivers.LoadStations("testdata/stations.json")

	if err != nil {
		t.Fatalf("can't read data file")
	}

	t.Run("Station number", func(t *testing.T) {
		wantStationsNumber := 444
		gotStationsNumber := len(data.Features)

		if gotStationsNumber != wantStationsNumber {
			t.Errorf("got: %d, want: %d", gotStationsNumber, wantStationsNumber)
		}
	})

	t.Run("Coordinate system", func(t *testing.T) {
		want := "EPSG:4326"

		got := data.Crs.Properties.Name

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

}
