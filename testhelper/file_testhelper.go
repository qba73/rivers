package testhelper

import (
	"os"
	"testing"
)

// TmpFile creates temporary file for testing.
func TmpFile(t *testing.T, dirname, filename string) *os.File {
	file, err := os.CreateTemp(dirname, filename)
	if err != nil {
		t.Fatalf("failure to create a temporary test file: %s", err)
	}
	return file
}

// TmpTextFile knows how to create temporary text file
// with provided content.
func TmpTextFile(t *testing.T, dirname, filename, content string) *os.File {
	file := TmpFile(t, dirname, filename)
	_, err := file.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write to temp file: %s", err)
	}
	return file
}

// TmpStationTestFile knows how to create a temporary file with content
// matching the latesttest.json file used as a data source for unit tests.
func TmpStationTestFile(t *testing.T, dirname, filename string) *os.File {
	file := TmpFile(t, dirname, filename)
	_, err := file.WriteString(stationsData)
	if err != nil {
		t.Fatalf("failed to write to temp file: %s", err)
	}
	return file
}

var stationsData = `{
	"type": "FeatureCollection",
	"crs": {
	  "type": "name",
	  "properties": { "name": "EPSG:4326" }
	},
	"features": [
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001041",
		  "station_name": "Sandy Mills",
		  "sensor_ref": "0001",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "1.715",
		  "err_code": 99,
		  "url": "/0000001041/0001/",
		  "csv_file": "/data/month/01041_0001.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.575758, 54.838318] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001041",
		  "station_name": "Sandy Mills",
		  "sensor_ref": "0002",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "4.800",
		  "err_code": 99,
		  "url": "/0000001041/0002/",
		  "csv_file": "/data/month/01041_0002.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.575758, 54.838318] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001041",
		  "station_name": "Sandy Mills",
		  "sensor_ref": "0003",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "13.000",
		  "err_code": 99,
		  "url": "/0000001041/0003/",
		  "csv_file": "/data/month/01041_0003.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.575758, 54.838318] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001041",
		  "station_name": "Sandy Mills",
		  "sensor_ref": "OD",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "8.060",
		  "err_code": 99,
		  "url": "/0000001041/OD/",
		  "csv_file": "/data/month/01041_OD.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.575758, 54.838318] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001043",
		  "station_name": "Ballybofey",
		  "sensor_ref": "0001",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "2.298",
		  "err_code": 99,
		  "url": "/0000001043/0001/",
		  "csv_file": "/data/month/01043_0001.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.790749, 54.799769] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001043",
		  "station_name": "Ballybofey",
		  "sensor_ref": "0002",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "5.470",
		  "err_code": 99,
		  "url": "/0000001043/0002/",
		  "csv_file": "/data/month/01043_0002.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.790749, 54.799769] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001043",
		  "station_name": "Ballybofey",
		  "sensor_ref": "0003",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "12.330",
		  "err_code": 99,
		  "url": "/0000001043/0003/",
		  "csv_file": "/data/month/01043_0003.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.790749, 54.799769] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000001043",
		  "station_name": "Ballybofey",
		  "sensor_ref": "OD",
		  "region_id": 3,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "18.228",
		  "err_code": 99,
		  "url": "/0000001043/OD/",
		  "csv_file": "/data/month/01043_OD.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.790749, 54.799769] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000003055",
		  "station_name": "Glaslough",
		  "sensor_ref": "0001",
		  "region_id": 10,
		  "datetime": "2021-02-18T05:00:00Z",
		  "value": "1.053",
		  "err_code": 99,
		  "url": "/0000003055/0001/",
		  "csv_file": "/data/month/03055_0001.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-6.894344, 54.323281] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000003055",
		  "station_name": "Glaslough",
		  "sensor_ref": "0002",
		  "region_id": 10,
		  "datetime": "2021-02-18T05:00:00Z",
		  "value": "6.300",
		  "err_code": 99,
		  "url": "/0000003055/0002/",
		  "csv_file": "/data/month/03055_0002.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-6.894344, 54.323281] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000003055",
		  "station_name": "Glaslough",
		  "sensor_ref": "0003",
		  "region_id": 10,
		  "datetime": "2021-02-18T05:00:00Z",
		  "value": "12.800",
		  "err_code": 99,
		  "url": "/0000003055/0003/",
		  "csv_file": "/data/month/03055_0003.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-6.894344, 54.323281] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000003055",
		  "station_name": "Glaslough",
		  "sensor_ref": "OD",
		  "region_id": 10,
		  "datetime": "2021-02-18T05:00:00Z",
		  "value": "36.840",
		  "err_code": 99,
		  "url": "/0000003055/OD/",
		  "csv_file": "/data/month/03055_OD.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-6.894344, 54.323281] }
	  },
	  {
		"type": "Feature",
		"properties": {
		  "station_ref": "0000003058",
		  "station_name": "Cappog Bridge",
		  "sensor_ref": "0001",
		  "region_id": 10,
		  "datetime": "2021-02-18T06:00:00Z",
		  "value": "1.233",
		  "err_code": 99,
		  "url": "/0000003058/0001/",
		  "csv_file": "/data/month/03058_0001.csv"
		},
		"geometry": { "type": "Point", "coordinates": [-7.021297, 54.266809] }
	  }
	]
}`
