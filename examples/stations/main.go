package main

import (
	"fmt"

	"github.com/qba73/rivers"
)

func main() {
	c := rivers.NewClient()

	stations, err := c.GetStations()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stations)
}
