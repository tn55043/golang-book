package main

import "fmt"

func main() {

	fmt.Printf("Convert FEET to METER Enter feet: ")
	var feet float32
	var meter float32
	fmt.Scanf("%f", &feet)
	meter=feet*0.3048
	fmt.Printf("meter : %.3f",meter)
}