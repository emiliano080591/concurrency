package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		var value = 97
		var result = 0
		goChan := make(chan int)
		mainChan := make(chan string)

		calculateSquare := func() {
			time.Sleep(time.Second * 3)
			result = value * value
			goChan <- result
		}

		reportResult := func() {
			time.Sleep(time.Second * 1)
			fmt.Println("The result of ", value, " squared is ", <-goChan)
			// blocks until it can read something from goChan - printed
			mainChan <- "You can quit now. I'm done"

		}

		go calculateSquare()
		go reportResult()

		<-mainChan
		wg.Done()
	}()
	wg.Wait()
}
