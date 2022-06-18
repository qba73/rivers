![Go](https://github.com/qba73/rivers/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/rivers)](https://goreportcard.com/report/github.com/qba73/rivers)
[![Maintainability](https://api.codeclimate.com/v1/badges/049487670cd44b2ab841/maintainability)](https://codeclimate.com/github/qba73/rivers/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/049487670cd44b2ab841/test_coverage)](https://codeclimate.com/github/qba73/rivers/test_coverage)

## rivers

```rivers``` is a Go library for reading water level and temperature data from stations located in rivers in Ireland. It allows to get data from more than 450 sensors located in 28 rivers.  

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

