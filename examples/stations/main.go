package main

import (
	"fmt"

	"github.com/qba73/rivers"
)

func main() {
	readings, err := rivers.GetLatestLevels()
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range readings {
		fmt.Println(r)
	}

	fmt.Println("=== Stations ===")

	fmt.Println(len(readings))
}
