package main

import (
	"fmt"
	"time"
)

func basicSelect() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		chan1 <- "Message from chan1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		chan2 <- "Message from chan1"
	}()

	// select waits for the first available channel
	select {
	case msg1 := <-chan1:
		fmt.Println("basicSelect:", msg1)
	case msg2 := <-chan2:
		fmt.Println("basicSelect:", msg2)
	}
}

func basicSelectWithDefault() {
	chan1 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		chan1 <- "Message from chan1"
	}()

	// select waits for the first available channel
	select {
	case msg1 := <-chan1:
		fmt.Println("basicSelectWithDefault:", msg1)
	default:
		fmt.Println("basicSelectWithDefault: No messages yet, skipping")
	}
}

func fanIn() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	chan3 := make(chan string)

	go func() {
		for i := 1; i <= 3; i++ {
			chan1 <- fmt.Sprintf("chan1 - msg %d", i)
		}
		fmt.Println("closing chan1")
		close(chan1)
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(200 * time.Millisecond)
			chan2 <- fmt.Sprintf("chan2 - msg %d", i)
		}
		fmt.Println("closing chan2")
		close(chan2)
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(50 * time.Millisecond)
			chan3 <- fmt.Sprintf("chan3 - msg %d", i)
		}
		fmt.Println("closing chan3")
		close(chan3)
	}()

	for {
		select {
		case msg, ok := <-chan1:
			if ok {
				fmt.Println("chan1: ", msg)
			} else {
				fmt.Println("chan1 is closed. Disabling chan1")
				chan1 = nil
			}
		case msg, ok := <-chan2:
			if ok {
				fmt.Println("chan2:", msg)
			} else {
				fmt.Println("chan2 is closed. Disabling chan2")
				chan2 = nil
			}
		case msg, ok := <-chan3:
			if ok {
				fmt.Println("chan3:", msg)
			} else {
				fmt.Println("chan3 is closed. Disabling chan2")
				chan3 = nil
			}
		}

		if chan1 == nil && chan2 == nil && chan3 == nil {
			break
		}
	}
}

func main() {
	basicSelect()
	basicSelectWithDefault()
	fanIn()
}
