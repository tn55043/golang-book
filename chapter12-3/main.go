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
	for i := 0; i <= 100; i++ {
		fmt.Println(i, fizzbuzz(i))
	}
}

func fizzbuzz(number int) string {
	fbTemplate := func(number int, str string) func(int) (string, bool) {
		return func(n int) (string, bool) {
			if n%number == 0 {
				return str, true
			}
			return "", false
		}
	}

	fbArrey := [...]func(n int) (string, bool){
		fbTemplate(15, "FizzBuzz"),
		fbTemplate(3, "Fizz"),
		fbTemplate(5, "Buzz"),
	}

	for i := 0; i < len(fbArrey); i++ {
		if str, ok := fbArrey[i](number); ok {
			return str
		}
	}
	return strconv.Itoa(number)
}
