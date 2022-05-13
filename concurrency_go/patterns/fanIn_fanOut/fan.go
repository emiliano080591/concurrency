package main

import (
	"fmt"
	"sync"
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

func merge(cs ...<-chan int) <-chan int {
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	// sincroniza
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := generator(2, 3, 4, 5, 6, 7)

	ch1 := square(in)
	ch2 := square(in)

	for n := range merge(ch1, ch2) {
		fmt.Println(n)
	}
}
