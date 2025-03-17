package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Println("value of i:", i)
	// p will store the memory address that points to i
	p := &i
	fmt.Println("memory address of i:", p)

	// you can read the value the pointer points to(i) by using *
	fmt.Println("value of i using *:", *p)

	// we can also change the value of i through *
	*p = 2
	fmt.Println("updated value of i using *:", i)
}