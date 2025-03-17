package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	pt := Vertex{2, 3}

	fmt.Println("vertex pt:", pt)
	fmt.Println("x value:", pt.X)
	fmt.Println("y value:", pt.Y)

	pt.X = 4
	fmt.Println("Updated x value:", pt.X)

	p := &pt
	fmt.Println("memory address of pt:", p)

	// we can update the value of pt using p like: (*pt).x but writing this is complex so we can just write:
	p.X = 5
	fmt.Println("Updated x value:", pt.X)

	fmt.Println()
	
	struct_literal()
}