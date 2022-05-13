package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// generator() -> square() ->
//								merge -> print
// 			   -> square() ->
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
				fmt.Println("Proccessing in generator num->", n)
				time.Sleep(time.Second)
			case <-done:
				return
			}

		}
	}()

	return out
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
				fmt.Println("Proccessing in square num->", n)
				time.Sleep(time.Second)
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	// sincroniza
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
			out <- n
		}
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
	done := make(chan struct{})
	in := generator(done, 2, 3, 4, 5, 6, 7)

	ch1 := square(done, in)
	ch2 := square(done, in)

	/*for n := range merge(ch1, ch2) {
		fmt.Println(n)
	}*/

	out := merge(done, ch1, ch2)
	fmt.Println(<-out)
	close(done)
	g := runtime.NumGoroutine()
	fmt.Println("Num goroutines->", g)
}
