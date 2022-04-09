package main

import (
	"fmt"

	"github.com/qba73/rivers"
)

func main() {
	readings, err := rivers.GetLatest()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(readings)
}
