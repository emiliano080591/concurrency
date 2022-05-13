package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	go doSomething()
	go doSomethingElse()

	time.Sleep(time.Second * 3)

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Println("Processes took ", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("\n I've done something")
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
}
