package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("recover")
		if p := recover(); p != nil {
			fmt.Printf("Panic found: %v\n", p)
		}
	}()

	fmt.Println("start main")
	panic("boom!")
}
