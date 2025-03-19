package main

import "fmt"

func declarations() {
	// way 1: creating a slice from an existing array
	nums := [5]int{1, 2, 3, 4, 5}
	var s1 []int = nums[1:4] // includes 1 and excludes 4
	fmt.Println("s1:", s1)

	// way 2: this way creates an array and returns a slice reference
	var s2 = []int{1, 2, 3}
	fmt.Println("s2:", s2)

	// way 3: shorthand way of creating slice
	s3 := []int{4, 5, 6}
	fmt.Println("s3:", s3)

	// way 4: to create a slice containing all the elements of the array
	s4 := nums[:]
	fmt.Println("s4 containing all the elements of nums array:", s4)

	// way 5: using make fn make([]T, len, cap) default value is 0
	s5 := make([]int, 3, 5)
	fmt.Println("s5:", s5)

	
}
