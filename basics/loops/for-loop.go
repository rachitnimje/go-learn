package main

import (
	"fmt"
)

func basic_for(n int) int {
	sum := 0
	for i := 1; i<=n; i++ {
		sum += i
	}
	return sum
}

func only_cond(n int) int {
	sum := 1
	for ; sum<n; {
		sum += sum
	}
	return sum
}

func while_loop(n int) int {
	sum := 1
	for sum<n {
		sum += sum
	}
	return sum
}

// func infinite_loop(n int) int {
// 	for {
// 		n += n
// 	}
// 	return n
// }

func main() {
	fmt.Println(basic_for(10))
	fmt.Println(only_cond(10))
	fmt.Println(while_loop(10))
}