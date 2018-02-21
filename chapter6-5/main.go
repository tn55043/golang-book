package main

import (
	"fmt"
	"strconv"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(fizzbuzz(number))
	}
}

/*
func fizzbuzz(number int) (res string) {
	if number%15 == 0 {
		res = "FizzBuzz"
	} else if number%3 == 0 {
		res = "Fizz"
	} else if number%5 == 0 {
		res = "Buzz"
	} else {
		res = strconv.Itoa(number)
	}
	return res
}
*/

func fizzbuzz(number int) string {
	ln := [3]int{15, 3, 5}
	str := [3]string{"FizzBuzz", "Fizz", "Buzz"}
	for i := 0; i < len(ln); i++ {
		if number%ln[i] == 0 {
			return str[i]
		}
	}
	return strconv.Itoa(number)
}
