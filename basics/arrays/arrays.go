package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := range len(arr1) {
		arr1[i] = i
	}

	var arr2 = [5]int{1, 3, 5, 7, 9}
	arr3 := [5]int {2, 4, 6, 8, 10}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

// var arr [5]int -> just initialization
// var arr = [3]int {1, 2, 3}
// arr := [3]int {1, 2, 3}