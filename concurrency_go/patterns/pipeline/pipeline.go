package main

import (
	"fmt"
	"time"
)

// generator() -> square() -> print
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Println("Proccessing in generator num->", n)
			time.Sleep(time.Second)
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			fmt.Println("Proccessing in square num->", n)
			time.Sleep(time.Second)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	//ch := generator(2, 3, 4, 5, 6, 7)
	//out := square(ch)

	for n := range square(generator(2, 3, 4, 5, 6, 7)) {
		fmt.Println(n)
	}
}
