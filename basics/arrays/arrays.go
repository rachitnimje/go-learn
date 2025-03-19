package main

import "fmt"

func main() {
	// way 1: just declaring
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println("arr1:", arr1)

	// way 2: declaring and initializing
	var arr2 = [3]int{
		1, 3, 5}
	fmt.Println("arr2:", arr2)
	
	// way 3: shorthand of way 2
	arr3 := [3]int {2, 4, 6}
	fmt.Println("arr3:", arr3)

	// way 4: compiler determines the length
	arr4 := [...]int {1, 2, 3}
	fmt.Println("arr4:", arr4)

	// properties
	fmt.Println("\nlength of arr1:", len(arr1))
	fmt.Println("element at position 2 in arr2:", arr2[1])

	// copy the array
	arr4 = arr2
	fmt.Println("updated arr4 using arr4 = arr2:", arr4)
	
	// makes a copy of arr3 into arr5 by value and not by reference
	arr5 := arr3
	fmt.Println("arr5:", arr5)

	arr5[1] = 10
	fmt.Println("updated arr5:", arr5)
	fmt.Println("arr3 remains as it is:", arr3)

	// passes the arr3 by value and changes the element in the array to check if the original chnages or not
	changeElement(arr3)
	fmt.Println("arr3 remains unchanged thus value is passed to the func and not reference", arr3)

	fmt.Println("iterating on arr using for loop (range):")
	for i, v := range arr3 {
		fmt.Println("index:", i, " value:", v)
	}
}

func changeElement(a [3]int) {
	a[2] = 12
	fmt.Println("method called, updated arr:", a)
}
