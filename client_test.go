package rivers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"
)

func startServer(path string, datafile string, t *testing.T) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//if r.URL.Path != path {
		//	t.Fatalf("incorrect URL: got %q", r.URL.Path)
		//}

		f, err := os.Open(datafile)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		io.Copy(rw, f)
	}))

	return ts
}

func TestGetStations(t *testing.T) {
	t.Parallel()

	ts := startServer("/geojson/", "testdata/stations_short.json", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	got, err := client.GetStations()
	if err != nil {
		t.Fatal(err)
	}

	want := rivers.Stations{
		Type: "FeatureCollection",
		Crs: rivers.Crs{
			Type: "name",
			Properties: rivers.CrsProperty{
				Name: "EPSG:4326"},
		},
		Features: []rivers.Feature{
			{
				Type: "Feature",
				Properties: rivers.Property{
					Name: "Sandy Mills",
					Ref:  "0000001041",
				},
				Geometry: rivers.Geometry{
					Type:        "Point",
					Coordinates: []float64{-7.575758, 54.838318},
				},
			},
			{
				Type: "Feature",
				Properties: rivers.Property{
					Name: "Ballybofey",
					Ref:  "0000001043",
				},
				Geometry: rivers.Geometry{
					Type:        "Point",
					Coordinates: []float64{-7.790749, 54.799769},
				},
			},
		},
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetLatest(t *testing.T) {
	t.Parallel()

	ts := startServer("/geojson/latest", "testdata/latest_short.json", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	got, err := client.GetLatest()
	if err != nil {
		t.Fatalf("GetLatest() got error %v", err)
	}

	want := rivers.StationsLatest{
		Type: "FeatureCollection",
		Crs: rivers.Crs{
			Type: "name",
			Properties: rivers.CrsProperty{
				Name: "EPSG:4326",
			},
		},
		Features: []rivers.FeatureLatest{
			{
				Type: "Feature",
				Properties: rivers.PropertyLatest{
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
				Geometry: rivers.Geometry{
					Type:        "Point",
					Coordinates: []float64{-7.575758, 54.838318},
				},
			},
			{
				Type: "Feature",
				Properties: rivers.PropertyLatest{
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
				Geometry: rivers.Geometry{
					Type:        "Point",
					Coordinates: []float64{-7.575758, 54.838318},
				},
			},
			{
				Type: "Feature",
				Properties: rivers.PropertyLatest{
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
				Geometry: rivers.Geometry{
					Type:        "Point",
					Coordinates: []float64{-7.575758, 54.838318},
				},
			},
			{
				Type: "Feature",
				Properties: rivers.PropertyLatest{
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
				Geometry: rivers.Geometry{
					Type:        "Point",
					Coordinates: []float64{-7.575758, 54.838318},
				},
			},
		},
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetDayLevel(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/day", "testdata/day_01041_0001.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.SensorReading{
		{
			time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			0.294,
		},
		{
			time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			0.293,
		},
		{
			time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			0.293,
		},
		{
			time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			0.293,
		},
	}

	stationID := "010104"
	got, err := client.GetDayLevel(stationID)
	if err != nil {
		t.Fatalf("client.GetDayLevel(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetWeekLevel(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/week", "testdata/week_01041_0001.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.SensorReading{
		{
			time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			0.294,
		},
		{
			time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			0.293,
		},
		{
			time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			0.293,
		},
		{
			time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			0.293,
		},
	}

	stationID := "010104"
	got, err := client.GetWeekLevel(stationID)
	if err != nil {
		t.Fatalf("client.GetWeekLevel(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetMonthLevel(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/month", "testdata/month_01041_0001.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.SensorReading{
		{
			time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			0.294,
		},
		{
			time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			0.293,
		},
		{
			time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			0.293,
		},
		{
			time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			0.293,
		},
	}

	stationID := "010104"
	got, err := client.GetMonthLevel(stationID)
	if err != nil {
		t.Fatalf("client.GetMonthLevel(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetDayTemperature(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/day", "testdata/day_01041_0002.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.SensorReading{
		{
			time.Date(2021, 07, 15, 22, 00, 00, 00, time.UTC),
			19.900,
		},
		{
			time.Date(2021, 07, 15, 23, 00, 00, 00, time.UTC),
			19.700,
		},
		{
			time.Date(2021, 07, 16, 00, 00, 00, 00, time.UTC),
			19.400,
		},
	}

	stationID := "010104"
	got, err := client.GetDayTemperature(stationID)
	if err != nil {
		t.Fatalf("GetDayTemperature(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetWeekTemperature(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/week", "testdata/week_01041_0002.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.SensorReading{
		{
			time.Date(2021, 07, 15, 22, 00, 00, 00, time.UTC),
			19.900,
		},
		{
			time.Date(2021, 07, 15, 23, 00, 00, 00, time.UTC),
			19.700,
		},
		{
			time.Date(2021, 07, 16, 00, 00, 00, 00, time.UTC),
			19.400,
		},
	}

	stationID := "010104"
	got, err := client.GetWeekTemperature(stationID)
	if err != nil {
		t.Fatalf("GetWeekTemperature(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetMonthTemperature(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/month", "testdata/month_01041_0002.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.SensorReading{
		{
			time.Date(2021, 07, 15, 22, 00, 00, 00, time.UTC),
			19.900,
		},
		{
			time.Date(2021, 07, 15, 23, 00, 00, 00, time.UTC),
			19.700,
		},
		{
			time.Date(2021, 07, 16, 00, 00, 00, 00, time.UTC),
			19.400,
		},
	}

	stationID := "010104"
	got, err := client.GetMonthTemperature(stationID)
	if err != nil {
		t.Fatalf("GetMonthTemperature(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
