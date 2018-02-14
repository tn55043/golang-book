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
