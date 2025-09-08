package main

import (
	"fmt"
	"sync"
)

func printMessageUsingChannels(ch chan string, wg *sync.WaitGroup, i int) {
	// using defer since this goroutine should be marked as done even if panic occurs
	defer wg.Done()
	ch <- fmt.Sprintf("Hello for the %d th time", i)
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		// waitGroup should be added to before starting the goroutine
		wg.Add(1)
		go printMessageUsingChannels(ch, &wg, i)
	}

	fmt.Println("Hello from the main function")

	// receives the values in another goroutines
	// add receiver to wait group too since receiver might still be running when the main exits
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}
		close(ch) // no more values will be sent to this channel, but the buffered values can be read/received
	}()

	// blocks the main goroutine till the counter reaches 0
	wg.Wait()
}
