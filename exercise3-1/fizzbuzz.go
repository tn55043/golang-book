package test_fizzbuzz

import "strconv"

func fizzbuzz(number int) string {
    var str string
    if number<1 {
        str = ""
    } else if number>100 {
        str = ""
    } else if number%15 == 0 {
        str = "FizzBuzz"
    } else if number%3 == 0 {
        str = "Fizz"
    } else if number%5 == 0 {
        str = "Buzz"
    } else {
        str = strconv.Itoa(number)
    }
    return str
}
