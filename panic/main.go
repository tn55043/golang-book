package main

import (
	"fmt"
)

func f(ch chan int) {
	defer func() {
		fmt.Println("Deferred by f")
	}()
	g()
	ch <- 0
}

func g() {
	defer func() {
		fmt.Println("Deferred by g")
	}()
	h()
}

func h() {
	defer func() {
		fmt.Println("Deferred by h")
	}()
	panic("boom!")
}

func main() {
	defer func() {
		fmt.Println("recover")
		if p := recover(); p != nil {
			fmt.Printf("Panic found: %v\n", p)
		}
	}()
	ch := make(chan int)
	go f(ch)
	<-ch
}

/*
func f() {
	defer func() {
		fmt.Println("Deferred by f")
	}()
	g()
}

func g() {
	defer func() {
		fmt.Println("Deferred by g")
	}()
	h()
}

func h() {
	defer func() {
		fmt.Println("Deferred by h")
	}()
	panic("boom!")
}

func main() {
	defer func() {
		fmt.Println("recover")
		if p := recover(); p != nil {
			fmt.Printf("Panic found: %v\n", p)
		}
	}()
	f()
}
*/
/*
func main() {
	defer func() {
		fmt.Println("recover")
		if p := recover(); p != nil {
			fmt.Printf("Panic found: %v\n", p)
		}
	}()

	fmt.Println("start main")
	panic("boom!")
}
*/
