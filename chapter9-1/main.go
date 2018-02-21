package main

import (
	"fmt"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	var c Circle
	fmt.Printf("c type: %T\n", c)
	fmt.Println(c.x, c.y, c.r)
	
	c1 := new(Circle)
	fmt.Printf("c1 type: %T\n", c1)
	fmt.Println(c1.x, c1.y, c1.r)
}
