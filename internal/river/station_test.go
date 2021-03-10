package river_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers/internal/river"
)

/*
func setupStationFile(t *testing.T) *os.File {
	return testhelper.TmpStationTestFile(t, ".", "st.json")
}

func cleanupStationsFile(file *os.File) {
	os.Remove(file.Name())
}
*/

var testFeatures = []river.Feature{
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001041",
			StationName: "Sandy Mills",
			SensorRef:   "0001",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "1.715",
			ErrCode:     99,
			URL:         "/0000001041/0001/",
			CSVFile:     "/data/month/01041_0001.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.575758, 54.838318}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001041",
			StationName: "Sandy Mills",
			SensorRef:   "0002",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "4.800",
			ErrCode:     99,
			URL:         "/0000001041/0002/",
			CSVFile:     "/data/month/01041_0002.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.575758, 54.838318}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001041",
			StationName: "Sandy Mills",
			SensorRef:   "0003",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "13.000",
			ErrCode:     99,
			URL:         "/0000001041/0003/",
			CSVFile:     "/data/month/01041_0003.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.575758, 54.838318}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001041",
			StationName: "Sandy Mills",
			SensorRef:   "OD",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "8.060",
			ErrCode:     99,
			URL:         "/0000001041/OD/",
			CSVFile:     "/data/month/01041_OD.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.575758, 54.838318}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001043",
			StationName: "Ballybofey",
			SensorRef:   "0001",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "2.298",
			ErrCode:     99,
			URL:         "/0000001043/0001/",
			CSVFile:     "/data/month/01043_0001.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.790749, 54.799769}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001043",
			StationName: "Ballybofey",
			SensorRef:   "0002",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "5.470",
			ErrCode:     99,
			URL:         "/0000001043/0002/",
			CSVFile:     "/data/month/01043_0002.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.790749, 54.799769}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001043",
			StationName: "Ballybofey",
			SensorRef:   "0003",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "12.330",
			ErrCode:     99,
			URL:         "/0000001043/0003/",
			CSVFile:     "/data/month/01043_0003.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.790749, 54.799769}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000001043",
			StationName: "Ballybofey",
			SensorRef:   "OD",
			RegionID:    3,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "18.228",
			ErrCode:     99,
			URL:         "/0000001043/OD/",
			CSVFile:     "/data/month/01043_OD.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.790749, 54.799769}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000003055",
			StationName: "Glaslough",
			SensorRef:   "0001",
			RegionID:    10,
			Timestamp:   "2021-02-18T05:00:00Z",
			Value:       "1.053",
			ErrCode:     99,
			URL:         "/0000003055/0001/",
			CSVFile:     "/data/month/03055_0001.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000003055",
			StationName: "Glaslough",
			SensorRef:   "0002",
			RegionID:    10,
			Timestamp:   "2021-02-18T05:00:00Z",
			Value:       "6.300",
			ErrCode:     99,
			URL:         "/0000003055/0002/",
			CSVFile:     "/data/month/03055_0002.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000003055",
			StationName: "Glaslough",
			SensorRef:   "0003",
			RegionID:    10,
			Timestamp:   "2021-02-18T05:00:00Z",
			Value:       "12.800",
			ErrCode:     99,
			URL:         "/0000003055/0003/",
			CSVFile:     "/data/month/03055_0003.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000003055",
			StationName: "Glaslough",
			SensorRef:   "OD",
			RegionID:    10,
			Timestamp:   "2021-02-18T05:00:00Z",
			Value:       "36.840",
			ErrCode:     99,
			URL:         "/0000003055/OD/",
			CSVFile:     "/data/month/03055_OD.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
	},
	{
		Type: "Feature",
		Properties: river.Property{
			StationRef:  "0000003058",
			StationName: "Cappog Bridge",
			SensorRef:   "0001",
			RegionID:    10,
			Timestamp:   "2021-02-18T06:00:00Z",
			Value:       "1.233",
			ErrCode:     99,
			URL:         "/0000003058/0001/",
			CSVFile:     "/data/month/03058_0001.csv",
		},
		Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.021297, 54.266809}},
	},
}

func TestLoadJSON(t *testing.T) {
	dataFile := "testdata/latesttest.json"
	st, err := river.LoadStations(dataFile)
	if err != nil {
		t.Fatalf("can't read data file: %s", err)
	}

	t.Run("Station number", func(t *testing.T) {
		wantStationsNumber := 13
		gotStationsNumber := len(st.Features)

		if gotStationsNumber != wantStationsNumber {
			t.Errorf("got: %d, want: %d", gotStationsNumber, wantStationsNumber)
		}
	})

	t.Run("Coordinate system", func(t *testing.T) {
		want := "EPSG:4326"

		got := st.Crs.Properties.Name

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestReadStations(t *testing.T) {
	f, err := os.Open("testdata/stationstest.json")
	if err != nil {
		t.Fatalf("error opening temp test file: %s", err)
	}
	defer f.Close()

	s, err := river.ReadStations(f)
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
	testFile := "testdata/latesttest.json"
	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("error opening temp test file: %s", err)
	}
	defer f.Close()

	s, err := river.ReadStations(f)
	if err != nil {
		t.Fatalf("error when reading stations: %s", err)
	}

	t.Run("Retrieve all features", func(t *testing.T) {
		got := s.GetAll()

		wantLen := 13
		if len(got.Features) != wantLen {
			t.Errorf("GetAll() got: %d, want: %d", len(got.Features), wantLen)
		}

		wantFeatures := testFeatures
		if !cmp.Equal(got.Features, wantFeatures) {
			t.Errorf("GetAll() got error \n%s", cmp.Diff(got.Features, wantFeatures))
		}
	})

	t.Run("Retrieve feature by name", func(t *testing.T) {
		station := "Glaslough"
		got := s.GetByName(station)

		wantLen := 4
		gotLen := len(got.Features)

		if gotLen != wantLen {
			t.Errorf("GetByName() got number of items: %d, want: %d", gotLen, wantLen)
		}

		wantFeatures := []river.Feature{
			{
				Type: "Feature",
				Properties: river.Property{
					StationRef:  "0000003055",
					StationName: "Glaslough",
					SensorRef:   "0001",
					RegionID:    10,
					Timestamp:   "2021-02-18T05:00:00Z",
					Value:       "1.053",
					ErrCode:     99,
					URL:         "/0000003055/0001/",
					CSVFile:     "/data/month/03055_0001.csv",
				},
				Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
			},
			{
				Type: "Feature",
				Properties: river.Property{
					StationRef:  "0000003055",
					StationName: "Glaslough",
					SensorRef:   "0002",
					RegionID:    10,
					Timestamp:   "2021-02-18T05:00:00Z",
					Value:       "6.300",
					ErrCode:     99,
					URL:         "/0000003055/0002/",
					CSVFile:     "/data/month/03055_0002.csv",
				},
				Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
			},
			{
				Type: "Feature",
				Properties: river.Property{
					StationRef:  "0000003055",
					StationName: "Glaslough",
					SensorRef:   "0003",
					RegionID:    10,
					Timestamp:   "2021-02-18T05:00:00Z",
					Value:       "12.800",
					ErrCode:     99,
					URL:         "/0000003055/0003/",
					CSVFile:     "/data/month/03055_0003.csv",
				},
				Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
			},
			{
				Type: "Feature",
				Properties: river.Property{
					StationRef:  "0000003055",
					StationName: "Glaslough",
					SensorRef:   "OD",
					RegionID:    10,
					Timestamp:   "2021-02-18T05:00:00Z",
					Value:       "36.840",
					ErrCode:     99,
					URL:         "/0000003055/OD/",
					CSVFile:     "/data/month/03055_OD.csv",
				},
				Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-6.894344, 54.323281}},
			},
		}

		if !cmp.Equal(got.Features, wantFeatures) {
			t.Errorf("GetByName(%s) \n%s", station, cmp.Diff(got.Features, wantFeatures))
		}
	})

	t.Run("Get feature by station ref number single", func(t *testing.T) {
		ref := "0000003058"
		got := s.GetByStationRef(ref)

		wantLen := 1
		gotLen := len(got.Features)

		if gotLen != wantLen {
			t.Errorf("GetByStationRef(%s) got number of features: %d, want: %d", ref, gotLen, wantLen)
		}

		wantFeatures := []river.Feature{
			{
				Type: "Feature",
				Properties: river.Property{
					StationRef:  "0000003058",
					StationName: "Cappog Bridge",
					SensorRef:   "0001",
					RegionID:    10,
					Timestamp:   "2021-02-18T06:00:00Z",
					Value:       "1.233",
					ErrCode:     99,
					URL:         "/0000003058/0001/",
					CSVFile:     "/data/month/03058_0001.csv",
				},
				Geometry: river.Geometry{Type: "Point", Coordinates: []float64{-7.021297, 54.266809}},
			},
		}

		if !cmp.Equal(got.Features, wantFeatures) {
			t.Errorf("GetByStationRef(%s) \n%s", ref, cmp.Diff(got.Features, wantFeatures))
		}
	})

	t.Run("Get feture by station ref number multiple", func(t *testing.T) {
		ref := "0000001043"
		got := s.GetByStationRef(ref)

		wantLen := 4
		gotLen := len(got.Features)

		if gotLen != wantLen {
			t.Errorf("GetByStationRef(%s) got number of features: %d, want: %d", ref, gotLen, wantLen)
		}
	})
}

/*
func TestStationsGetBySensorRef(t *testing.T) {
	t.Parallel()

	fl := setupStationFile(t)
	defer cleanupStationsFile(fl)

	f, err := os.Open(fl.Name())
	if err != nil {
		t.Errorf("error creating file: %s", err)
	}

	s, err := river.ReadStations(f)
	if err != nil {
		t.Fatalf("error when reading stations: %s", err)
	}

	got := s.GetBySensorRef("0003")

	wantLen := 1
	gotLen := len(got.Features)

	if gotLen != wantLen {
		t.Errorf("GetBySensorRef(%s), got: %d, want: %d", "0003", gotLen, wantLen)
	}

}
*/
