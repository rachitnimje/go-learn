package main

import "fmt"

func properties() {
	// slice don't actually store any data
	// they are like references to the array or part of array
	names := [4]string{"Rachit", "Tanisha", "Dhruvesh", "Tonya"}
	f1 := names[0:2]
	fmt.Println("\nnames:", names, "\nf1:", f1)

	f1[1] = "Chingus"
	fmt.Println("Updated f1:", f1)
	fmt.Println("Updated names:", names)

	// length of slice is the number of elements in the slice
	fmt.Println("\nlength of slice s4:", len(f1))
	// capacity of slice is the number of elements in the underlying array from which the slice was created
	fmt.Println("capacity of slice s4:", cap(f1))

	// re-slicing the slice f1
	f1 = f1[:cap(names)]
	fmt.Println("\nlength of slice after re-slicing s4:", len(f1))
	fmt.Println("capacity of slice after re-slicing s4:", cap(f1))

	// on appending a slice, the original array remains unchanged 
	// new array is created with the same values plus the appended value with the capacity of twice as much as the original
	// and the reference of the newly created array is returned to the original slice 
	f1 = append(f1, "Akbar")
	fmt.Println("\nAfter appending on f1 slice, f1:", f1)
	fmt.Println("After appending on f1 slice, names:", names)

	// we can see that the array remains unchanged on changing the elements of the slice
	f1[2] = "Chicken"
	fmt.Println("\nAfter Changing the value of slice f1 after appending:", f1)
	fmt.Println("After Changing the value of slice f1 after appending:", names)

}
