package main

import (
	"fmt"
)

func main() {
	var addVar func(int, int) int
	addVar = func(a, b int) int {
		return a + b
	}
	// first class function
	fmt.Println(addVar(2, 3))

	fmt.Println(add(2, 3))
}

func add(a, b int) int {
	return a + b
}