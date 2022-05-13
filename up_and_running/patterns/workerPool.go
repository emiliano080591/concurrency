package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 10
	jobsChan := make(chan int, numJobs)
	completedJobsChan := make(chan int, numJobs)

	for i := 1; i <= 3; i++ { //this is the number of workers
		go worker(i, jobsChan, completedJobsChan)
	}

	for j := 1; j <= numJobs; j++ {
		// Se llena jobsChan para que los workers trabajen
		jobsChan <- j // This loads the jobsChan channel with job numbers
	}
	close(jobsChan) // Close jobsChan channel for input after all jobs have been loaded.  Channel must be closed in order to call "range" function.

	for a := 1; a <= numJobs; a++ {
		// espera a que terminen todos lo jobs
		<-completedJobsChan // Reads the completedJobsChan channel and does nothing with the contents.  Point is to clear the channel and to delay termination of the program until all jobs are reported as finished.
	}
}

func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) {
	for j := range jobsChan { // iterates (and RECEIVES) each and all the jobs in the channel.  Interesting that range seems to have its own receiver.
		fmt.Println("worker", id, "started  job", j, "with", len(jobsChan), "jobs left to process")
		time.Sleep(time.Second * 2) // simulates "work" that takes sleep time to complete
		fmt.Println("worker", id, "             finished job", j)
		completedJobsChan <- j // Loads finished job numbers into the completedJobsChan channel.
	}
}
