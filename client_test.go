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
		f, err := os.Open(datafile)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		io.Copy(rw, f)
	}))
	return ts
}

func TestRiversClient_GetsLatestWaterLevelReadings(t *testing.T) {
	t.Parallel()
	ts := startServer("/geojson/latest", "testdata/latest_short.json", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	got, err := client.GetLatestWaterLevels()
	if err != nil {
		t.Fatalf("GetLatest() got error %v", err)
	}

	want := []rivers.StationWaterLevelReading{
		{
			StationID:  "0000001041",
			Readtime:   time.Date(2021, 02, 18, 06, 00, 00, 00, time.UTC),
			WaterLevel: 1.715,
		},
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_GetsDayWaterLevels(t *testing.T) {
	t.Parallel()
	ts := startServer("/data/day", "testdata/day_01041_0001.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.Reading{
		{
			Timestamp: time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			Value:     0.294,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			Value:     0.293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			Value:     0.293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			Value:     0.293,
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

func TestRiversClient_GetsWeekWaterLevels(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/week", "testdata/week_01041_0001.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.Reading{
		{
			Timestamp: time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			Value:     0.294,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			Value:     0.293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			Value:     0.293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			Value:     0.293,
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

func TestRiversClient_GetsMonthWaterLevel(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/month", "testdata/month_01041_0001.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.Reading{
		{
			Timestamp: time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			Value:     0.294,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			Value:     0.293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			Value:     0.293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			Value:     0.293,
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

func TestRiversClient_GetsDayWaterTemperature(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/day", "testdata/day_01041_0002.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.Reading{
		{
			Timestamp: time.Date(2021, 07, 15, 22, 00, 00, 00, time.UTC),
			Value:     19.900,
		},
		{
			Timestamp: time.Date(2021, 07, 15, 23, 00, 00, 00, time.UTC),
			Value:     19.700,
		},
		{
			Timestamp: time.Date(2021, 07, 16, 00, 00, 00, 00, time.UTC),
			Value:     19.400,
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

func TestRiversClient_GetsWeekWaterTemperature(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/week", "testdata/week_01041_0002.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.Reading{
		{
			Timestamp: time.Date(2021, 07, 15, 22, 00, 00, 00, time.UTC),
			Value:     19.900,
		},
		{
			Timestamp: time.Date(2021, 07, 15, 23, 00, 00, 00, time.UTC),
			Value:     19.700,
		},
		{
			Timestamp: time.Date(2021, 07, 16, 00, 00, 00, 00, time.UTC),
			Value:     19.400,
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

func TestRiversClient_GetsMonthWaterTemperature(t *testing.T) {
	t.Parallel()

	ts := startServer("/data/month", "testdata/month_01041_0002.csv", t)
	defer ts.Close()

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.Reading{
		{
			Timestamp: time.Date(2021, 07, 15, 22, 00, 00, 00, time.UTC),
			Value:     19.900,
		},
		{
			Timestamp: time.Date(2021, 07, 15, 23, 00, 00, 00, time.UTC),
			Value:     19.700,
		},
		{
			Timestamp: time.Date(2021, 07, 16, 00, 00, 00, 00, time.UTC),
			Value:     19.400,
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
