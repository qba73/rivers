package main

import (
	"fmt"
	"log"

	"github.com/qba73/rivers"
)

func main() {
	readings, err := rivers.GetLatestWaterLevels()
	if err != nil {
		log.Println(err)
	}
	for _, r := range readings {
		fmt.Printf("Station: %s, ID: %s, RegionID: %d, Time: %s, Water level: %.2f\n", r.Name, r.StationID, r.RegionID, r.Readtime, r.WaterLevel)
	}
}
