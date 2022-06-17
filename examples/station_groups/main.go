package main

import (
	"fmt"
	"log"

	"github.com/qba73/rivers"
)

func main() {
	client := rivers.NewClient()

	// groupID indicates which group station readings to retrieve.
	// groupID value is between 1 and 28.
	stations, err := client.GetStationGroupTemperature(1)
	if err != nil {
		log.Fatalln(err)
	}

	for _, station := range stations {
		fmt.Printf("%v\n", station)
	}
}
