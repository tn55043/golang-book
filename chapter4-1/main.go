package main

import (
	"fmt"
)

func main() {
	// Long Declaration
	var a string = "Hello, World"
	fmt.Println(a)

	var b string
	b = "Hello, World"
	fmt.Println(b)

	//Short Declaration
	// Type Inference
	c := "Hello, World"
	fmt.Println(c)
	fmt.Printf("Type: %T\n", c)
	d := 123
	fmt.Println(d)
	fmt.Printf("Type: %T\n", d)
	e := 12.12
	fmt.Println(e)
	fmt.Printf("Type: %T\n", e)

	const f string = "Constant Var"

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	l1, l2 := "first", "second"
	fmt.Println(l1)
	fmt.Println(l2)

	m1, m2 := "first", "second"
	m1, m2 = m2, m1
	fmt.Println(m1)
	fmt.Println(m2)
}

var (
	i = 5
	j = 10
	k = 15
)