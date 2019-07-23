package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v *Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, Abs(v))
	Scale(v, 5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, Abs(v))
}
