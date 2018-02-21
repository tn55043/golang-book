package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	double(slice)
	fmt.Printf("origin addr %p\n", slice)
	fmt.Printf("original %v\n", slice)
}

func double(nums []int) {
	fmt.Printf("double addr %p\n", nums)
	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}
	fmt.Println(nums)
}
