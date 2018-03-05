package main

import (
	"fmt"
	"strconv"
)

func main() {
	addFunc := func(a int) func(b int) int {
		return func(b int) int {
			return a + b
		}
	}
	addNewFunc := addFunc(2)
	fmt.Println(addNewFunc(3))

	// FizzBuzz Program
	for i := 0; i < 100; i++ {
		fmt.Println(i, fizzbuzz(i))
	}
}

func fizzbuzz(number int) string {
	fizzbuzzFunc := func(n int) (string, bool) {
		if n%15 == 0 {
			return "FizzBuzz", true
		}
		return "", false
	}

	fizzFunc := func(n int) (string, bool) {
		if n%3 == 0 {
			return "Fizz", true
		}
		return "", false
	}

	buzzFunc := func(n int) (string, bool) {
		if n%5 == 0 {
			return "Buzz", true
		}
		return "", false
	}

	fbArrey := [...]func(n int) (string, bool){
		fizzbuzzFunc,
		fizzFunc,
		buzzFunc,
	}

	for i := 0; i < len(fbArrey); i++ {
		if str, ok := fbArrey[i](number); ok {
			return str
		}
	}
	return strconv.Itoa(number)
}
