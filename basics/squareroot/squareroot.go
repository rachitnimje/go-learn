package main

import (
	"fmt"
	"math"
)

// the z² − x above is how far away z² is from where it needs to be (x), 
// the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. 
// This general approach is called Newton's method. 

func Sqrt(x float64) float64 {
	z := 1.0
	temp := 0.0

	for i:=0; i<10 && math.Abs(temp-z) > 0.000000001; i++ {
		fmt.Println(i, z)
		temp = z
		z -= (z*z - x) / (2*z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
