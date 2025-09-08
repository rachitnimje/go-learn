package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		results <- job * 2
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int)
	results := make(chan int)
	done := make(chan bool)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for a := 1; a <= numJobs; a++ {
			fmt.Println("Result:", <-results)
		}
		done <- true
	}()

	for job := 1; job <= numJobs; job++ {
		jobs <- job
	}
	close(jobs)

	<-done
}
