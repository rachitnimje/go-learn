package main

import (
	"fmt"
	"math"
)

func is_odd(n int) bool {
	if n % 2 == 0 {
		return false
	}
	return true
}

func is_greater_than_n_square(x, n float64) bool {
	// we can also write a short stmt in the if condition like below
	if m := math.Pow(n, 2); x > m {
		return true
	}  else {
		// we can also access the variable declared inside if statement in the else block
		fmt.Printf("%v <= %v\n", x, m)
	}
	return false
}

func main() {
	fmt.Println(is_odd(4))
	fmt.Println(is_greater_than_n_square(60, 8))
}