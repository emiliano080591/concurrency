package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(4)

	start := time.Now()
	counta()
	countb()
	countc()
	countd()
	countf()
	countg()
	counth()

	elpased := time.Since(start)
	fmt.Println("Processes took ", elpased)
}

func counth() {
	fmt.Println("HHHHH is starting ")
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Println("HHHHH is done! ")
}

func countg() {
	fmt.Println("GGGGGG is starting ")
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Println("GGGGGG is done! ")
}

func countf() {
	fmt.Println("FFFFFF is starting ")
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Println("FFFFFF is done! ")
}

func countd() {
	fmt.Println("DDDDDD is starting ")
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Println("DDDDDD is done! ")
}

func countc() {
	fmt.Println("CCCCC is starting ")
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Println("CCCCC is done! ")
}

func countb() {
	fmt.Println("BBBBB is starting ")
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Println("BBBBB is done! ")
}

func counta() {
	fmt.Println("AAAAA is starting ")
	for i := 0; i < 10_000_000_000; i++ {
	}
	fmt.Println("AAAAA is done! ")
}
