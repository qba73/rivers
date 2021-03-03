package rivers_test

import (
	"testing"

	"github.com/qba73/rivers"
)

func TestLoadStations(t *testing.T) {
	path := "testdata/stations.json"

	var stations rivers.Stations

	err := rivers.LoadStations(path, &stations)

	if err != nil {
		t.Fatalf("can't read data file")
	}

	t.Run("Station number", func(t *testing.T) {
		wantStationsNumber := 444
		gotStationsNumber := len(stations.Features)

		if gotStationsNumber != wantStationsNumber {
			t.Errorf("got: %d, want: %d", gotStationsNumber, wantStationsNumber)
		}
	})

	t.Run("Coordinate system", func(t *testing.T) {
		want := "EPSG:4326"

		got := stations.Crs.Properties.Name

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

}
