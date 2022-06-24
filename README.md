![Go](https://github.com/qba73/rivers/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/rivers)](https://goreportcard.com/report/github.com/qba73/rivers)
[![Maintainability](https://api.codeclimate.com/v1/badges/049487670cd44b2ab841/maintainability)](https://codeclimate.com/github/qba73/rivers/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/049487670cd44b2ab841/test_coverage)](https://codeclimate.com/github/qba73/rivers/test_coverage)

## rivers

```rivers``` is a Go library and a cli utility for reading water level and temperature data from stations located in rivers in Ireland. It allows to get data from more than 450 sensors located in 28 rivers.  

## Using the Go library
Import the library using:
```go
import "github.com/qba73/rivers" 
```
## Creating a client
Creat a new ```Client``` object by calling ```rivers.NewClient()```:
```go
client := rivers.NewClient()
```
## Retrieving latest water level redings

```go
client.GetLatestWaterLevels()
```
or
```go
rivers.GetLatestWaterLevels()
```
## A complete example program
You can see an example programs which retrieves water level data in the [examples/stations](examples/stations/main.go) folder.

## Bugs and feature requests
If you find a bug in the ```rivers``` client or library, please [open an issue](https://github.com/qba73/rivers/issues). Similarly, if you'd like a feature added or improved, let me know via an issue.

Not all the functionality of the [water level](https://waterlevel.ie) is implemented yet.

Pull requests welcome!

## Using the CLI utility for reading water levels

Get readings from stations:
```
$ waterlevel
time: 2022-06-24 05:45:00 +0000 UTC, station: Sandy Mills, id: 1041, level: 0.21
time: 2022-06-24 05:15:00 +0000 UTC, station: Ballybofey, id: 1043, level: 0.46
time: 2022-06-24 05:45:00 +0000 UTC, station: Glaslough, id: 3055, level: 0.43
time: 2022-06-24 05:45:00 +0000 UTC, station: Cappog Bridge, id: 3058, level: 0.20
time: 2022-06-24 05:45:00 +0000 UTC, station: Moyles Mill, id: 6011, level: 0.39
time: 2022-06-24 05:45:00 +0000 UTC, station: Clarebane, id: 6012, level: 0.24
time: 2022-06-24 05:45:00 +0000 UTC, station: Charleville Weir, id: 6013, level: 0.28
time: 2022-06-24 05:45:00 +0000 UTC, station: Tallanstown Weir, id: 6014, level: 0.29
time: 2022-06-24 05:45:00 +0000 UTC, station: Mansfieldstown, id: 6021, level: -0.07
[...]
time: 2022-06-24 05:45:00 +0000 UTC, station: Ballyloskey, id: 40008, level: 0.51
time: 2022-06-24 05:45:00 +0000 UTC, station: Malin Head, id: 40060, level: 3.50
```

Save readings to a file:
```
$ waterlevel > levels.txt
```

### Disclaimer

```rivers``` project processes *Irish Public Sector Information* licensed under a Creative Commons Attribution 4.0 International (CC BY 4.0), [licence](http://waterlevel.ie) provided by the Office of Public Works.
Data is licensed under the [Creative Commons By Attribution (CC-BY) version 4.0 license](https://creativecommons.org/licenses/by/4.0/legalcode) - see [summary](https://creativecommons.org/licenses/by/4.0/).
This is in line with Irish Government [PER Circular 12 of 2016](http://circulars.gov.ie/pdf/circular/per/2016/12.pdf) and policy on Open Data [data.gov.ie](https://data.gov.ie/data).
