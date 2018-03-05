package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(9)

	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(n, ":", i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished")
}
