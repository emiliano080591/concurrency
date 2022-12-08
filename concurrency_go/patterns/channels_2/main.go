package main

import "fmt"

func sendValues(myIntChannel chan int) {
	for i := 1; i < 10; i++ {
		myIntChannel <- i
	}
	close(myIntChannel)
}

func main() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	//for i := range myIntChannel {
	//	fmt.Println(i)
	//}

	for {
		select {
		case s1, ok := <-myIntChannel:
			if ok {
				fmt.Println(s1)
			} else {
				return
			}
		}

	}

}
