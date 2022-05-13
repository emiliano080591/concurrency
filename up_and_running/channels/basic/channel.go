package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {
	start := time.Now()

	go doSomething()
	go doSomethingElse()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Println("Processes took ", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("\n I've done something")
	ch <- "doSomething finish"
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
	ch <- "doSomething else finish"
}
