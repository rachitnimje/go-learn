package main

import (
	"fmt"
)

func ex1() {
	defer fmt.Println("hello")
	
	fmt.Println("world")
}

func ex2() {
	fmt.Println("Counting:")

	for i := 1; i<10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	ex1()
	ex2()
}