package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 1; i <= 6; i++ {
		realnum := 9

		/*
		fmt.Printf("Input your number: ")
		fmt.Scanf("%d\n", &num)
		*/
		
		num = randInt(1, 30)
		//fmt.Println(randomString(10))

		if i > 5 {
			fmt.Println("End")
		} else {
			if num < realnum {
				fmt.Println(num, ": More")
			} else if num > realnum {
				fmt.Println(num, ": Less")
			} else {
				fmt.Println(num, ": Found")
				i=6
			}
		}
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
