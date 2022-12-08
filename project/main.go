package main

import (
	"fmt"
	"github.com/emiliano080591/concurrency/project/repository"
	"log"
)

func main() {
	n := 200
	result, err := repository.Fetch(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", result.Title)
}
