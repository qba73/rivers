package rivers_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/rivers"
)

func TestRiversClient_GetsLatestWaterLevelReadings(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/geojson/latest", "testdata/latest_short.json", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	got, err := client.GetLatestWaterLevels(context.Background())
	if err != nil {
		t.Fatalf("GetLatest() got error %v", err)
	}

	want := []rivers.StationWaterLevelReading{
		{
			StationID:  1041,
			Name:       "Sandy Mills",
			Readtime:   time.Date(2021, 02, 18, 06, 00, 00, 00, time.UTC),
			WaterLevel: 1715,
		},
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_GetsDayWaterLevels(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/data/day", "testdata/day_01041_0001.csv", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.WaterLevelReading{
		{
			Timestamp: time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			Value:     294,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			Value:     293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			Value:     293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			Value:     293,
		},
	}

	stationID := "010104"
	got, err := client.GetDayLevel(context.Background(), stationID)
	if err != nil {
		t.Fatalf("client.GetDayLevel(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_GetsWeekWaterLevels(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/data/week", "testdata/week_01041_0001.csv", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.WaterLevelReading{
		{
			Timestamp: time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			Value:     294,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			Value:     293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			Value:     293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			Value:     293,
		},
	}

	stationID := "010104"
	got, err := client.GetWeekLevel(context.Background(), stationID)
	if err != nil {
		t.Fatalf("client.GetWeekLevel(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_GetsMonthWaterLevel(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/data/month", "testdata/month_01041_0001.csv", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.WaterLevelReading{
		{
			Timestamp: time.Date(2021, 07, 10, 00, 00, 00, 00, time.UTC),
			Value:     294,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 15, 00, 00, time.UTC),
			Value:     293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 30, 00, 00, time.UTC),
			Value:     293,
		},
		{
			Timestamp: time.Date(2021, 07, 10, 00, 45, 00, 00, time.UTC),
			Value:     293,
		},
	}

	stationID := "010104"
	got, err := client.GetMonthLevel(context.Background(), stationID)
	if err != nil {
		t.Fatalf("client.GetMonthLevel(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_GetsDayWaterTemperature(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/data/day", "testdata/day_01041_0002.csv", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.WaterTemperatureReading{
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
	got, err := client.GetDayTemperature(context.Background(), stationID)
	if err != nil {
		t.Fatalf("GetDayTemperature(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_GetsWeekWaterTemperature(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/data/week", "testdata/week_01041_0002.csv", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.WaterTemperatureReading{
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
	got, err := client.GetWeekTemperature(context.Background(), stationID)
	if err != nil {
		t.Fatalf("GetWeekTemperature(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_GetsMonthWaterTemperature(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/data/month", "testdata/month_01041_0002.csv", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.WaterTemperatureReading{
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
	got, err := client.GetMonthTemperature(context.Background(), stationID)
	if err != nil {
		t.Fatalf("GetMonthTemperature(%q) got error %v", stationID, err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRiversClient_RetrievesGroupWaterLevel(t *testing.T) {
	t.Parallel()
	ts := newTestServer("/data/group", "testdata/group_1.csv", t)

	client := rivers.NewClient()
	client.BaseURL = ts.URL

	want := []rivers.StationWaterLevelReading{
		{
			Name:       "John's Bridge Nore",
			Readtime:   time.Date(2021, 06, 15, 22, 00, 00, 00, time.UTC),
			WaterLevel: 466,
		},
		{
			Name:       "Dinin Bridge",
			Readtime:   time.Date(2021, 06, 15, 22, 00, 00, 00, time.UTC),
			WaterLevel: 53,
		},
		{
			Name:       "Brownsbarn",
			Readtime:   time.Date(2021, 06, 15, 22, 00, 00, 00, time.UTC),
			WaterLevel: 413,
		},
		{
			Name:       "Mount Juliet",
			Readtime:   time.Date(2021, 06, 15, 22, 00, 00, 00, time.UTC),
			WaterLevel: 451,
		},
	}

	groupID := 1
	got, err := client.GetGroupWaterLevel(context.Background(), groupID)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func newTestServer(path string, datafile string, t *testing.T) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		f, err := os.Open(datafile)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		io.Copy(rw, f)
	}))

	t.Cleanup(func() {
		ts.Close()
	})

	return ts
}
