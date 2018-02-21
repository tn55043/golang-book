package main

import (
	"fmt"
)

func main() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	double(m)
	fmt.Printf("origin addr %p\n", m)
	fmt.Printf("original %v\n", m)
}

func double(nums map[int]int) {
	fmt.Printf("double addr %p\n", nums)
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)
}
