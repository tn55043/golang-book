package main

import (
	"fmt"
	"time"
)

func main() {
	volumn := 200
	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func order(volumn int) (container []string) {
	for i := 0; i <= volumn; i++ {
		coffee := recieveOrder(i)
		go func() {
			coffee = brew(coffee)
			container = append(container, serve(coffee))
		}()
	}
	return
}

func recieveOrder(number int) string {
	// cashier recieve order
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("order: %d", number)
}

func brew(order string) string {
	// barista brew coffee
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", order, "espresso")
}

func serve(coffee string) string {
	// waiter serve coffee
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("%s %s", coffee, "ready :)")
}