package main

import "fmt"

func main() {

	fmt.Printf("Convert Farenheit to Celsius Enter temp(F): ")
	var Farenheit float32
	var Celsius float32
	fmt.Scanf("%f", &Farenheit)
	Celsius=((Farenheit-32)*5)/9
	fmt.Printf("Temp(C) : %.2f",Celsius)
}