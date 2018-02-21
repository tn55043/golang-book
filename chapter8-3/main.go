package main

import (
	"fmt"
)

func main() {
	arrey := [3]int{1, 2, 3}
	double(arrey)
	fmt.Printf("original addr %p\n", &arrey)
	fmt.Printf("original %v\n", arrey)
}

func double(nums [3]int) {
	fmt.Printf("double addr %p\n", &nums)
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)
}
