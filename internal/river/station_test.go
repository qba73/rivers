package river_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers/internal/river"
	"github.com/qba73/rivers/internal/river/testhelper"
)

func setupStations(t *testing.T) *os.File {
	return testhelper.TmpTextFile(t, "", "stations.json", `{"type": "FeatureCollection", 
	"crs": {"type": "name", "properties": {"name": "EPSG:4326"}}, 
	"features": [
		{"type": "Feature", "properties": {"name": "Sandy Mills", "ref": "0000001041"}, "geometry": {"type": "Point", "coordinates": [-7.575758, 54.838318]}},
		{"type": "Feature", "properties": {"name": "Ballybofey", "ref": "0000001043"}, "geometry": {"type": "Point", "coordinates": [-7.790749, 54.799769]}},
		{"type": "Feature", "properties": {"name": "Glaslough", "ref": "0000003055"}, "geometry": {"type": "Point", "coordinates": [-6.894344, 54.323281]}},
		{"type": "Feature", "properties": {"name": "Cappog Bridge", "ref": "0000003058"}, "geometry": {"type": "Point", "coordinates": [-7.021297, 54.266809]}},
		{"type": "Feature", "properties": {"name": "Moyles Mill", "ref": "0000006011"}, "geometry": {"type": "Point", "coordinates": [-6.596077, 54.011574]}},
		{"type": "Feature", "properties": {"name": "Clarebane", "ref": "0000006012"}, "geometry": {"type": "Point", "coordinates": [-6.666056, 54.092856]}}]}`)
}

func cleanupStations(file *os.File) {
	os.Remove(file.Name())
}

func TestLoadStations(t *testing.T) {
	path := "testdata/stations.json"

	var stations river.Stations

	err := river.LoadStations(path, &stations)

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

func TestReadStations(t *testing.T) {
	file := setupStations(t)
	defer cleanupStations(file)

	stationsFile, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("error opening temp test file: %s", err)
	}

	s, err := river.ReadStations(stationsFile)
	if err != nil {
		t.Fatalf("error when reading stations: %s", err)
	}

	t.Run("Coordinate system", func(t *testing.T) {
		want := "EPSG:4326"

		got := s.Crs.Properties.Name

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

}

func TestStations(t *testing.T) {
	file := setupStations(t)
	defer cleanupStations(file)

	stationsFile, err := os.Open(file.Name())
	if err != nil {
		t.Fatalf("error opening temp test file: %s", err)
	}

	s, err := river.ReadStations(stationsFile)
	if err != nil {
		t.Fatalf("error when reading stations: %s", err)
	}

	t.Run("Retrieve all features", func(t *testing.T) {
		got := s.GetAllFeatures()

		wantLen := 6
		if len(got) != wantLen {
			t.Errorf("GetAllFeatures() got: %d, want: %d", len(got), wantLen)
		}

		wantFeatures := []river.Feature{
			{Type: "Feature", Properties: river.Property{Name: "Sandy Mills", Ref: "0000001041"}, Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.575758, 54.838318}}},
			{Type: "Feature", Properties: river.Property{Name: "Ballybofey", Ref: "0000001043"}, Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.790749, 54.799769}}},
			{Type: "Feature", Properties: river.Property{Name: "Glaslough", Ref: "0000003055"}, Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}}},
			{Type: "Feature", Properties: river.Property{Name: "Cappog Bridge", Ref: "0000003058"}, Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.021297, 54.266809}}},
			{Type: "Feature", Properties: river.Property{Name: "Moyles Mill", Ref: "0000006011"}, Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.596077, 54.011574}}},
			{Type: "Feature", Properties: river.Property{Name: "Clarebane", Ref: "0000006012"}, Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.666056, 54.092856}}},
		}

		if !cmp.Equal(got, wantFeatures) {
			t.Errorf("GetAllFeatures() got error \n%s", cmp.Diff(got, wantFeatures))
		}
	})

	t.Run("Retrieve feature by name", func(t *testing.T) {
		station := "Glaslough"
		got := s.GetFeatureByName(station)

		wantFeature := river.Feature{Type: "Feature", Properties: river.Property{Name: "Glaslough", Ref: "0000003055"}, Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}}}

		if !cmp.Equal(got, wantFeature) {
			t.Errorf("GetFeatureByName(%s) \n%s", station, cmp.Diff(got, wantFeature))
		}
	})

}
