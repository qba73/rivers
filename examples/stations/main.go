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
		fmt.Printf("Station: %s, ID: %d, Time: %s, Water level: %.d\n", r.Name, r.StationID, r.Readtime, r.WaterLevel)
	}
}
