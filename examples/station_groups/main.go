package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qba73/rivers"
)

func main() {
	client := rivers.NewClient()

	// groupID indicates which group station readings to retrieve.
	// groupID value is between 1 and 28.
	stations, err := client.GetGroupWaterLevel(context.Background(), 1)
	if err != nil {
		log.Fatalln(err)
	}

	for _, station := range stations {
		fmt.Println(station)
	}
}
