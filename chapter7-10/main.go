package main

import (
	"fmt"
)

func main() {
	maps := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, number := range maps {
		fmt.Println(key, number)
	}
}
