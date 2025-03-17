package main

import "fmt"

func struct_literal() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}		// value of Y is set to 0
		v3 = Vertex{}			// value of both is set to 0
		v4 = &v1
		v5 = &Vertex{1, 2}
	)

	fmt.Println(v1, v2, v3, v4, v5)
}