package main

import (
	"fmt"
)

func main() {
	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	//Short Declaration
	// Type Inference
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)
	u := 123
	fmt.Println(u)
	fmt.Printf("Type: %T\n", u)
	v := 12.12
	fmt.Println(v)
	fmt.Printf("Type: %T\n", v)
}
