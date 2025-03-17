package main

import (
	"fmt"
	"runtime"
	"time"
)

func switch_1() {
	fmt.Print("Code is running on ")
	count := 0

	// in go, the break statement after each case is given by-default
	// like in if statements, we can declare var in switch statement
	switch os := runtime.GOOS; os {
	case "linux":
		count++
		fmt.Print("Linux.")
	case "macos" :
		count++
		fmt.Print("MacOS.")
	case "windows":
		count++
		fmt.Print("Windows.")
	default:
		count++
		fmt.Printf("%s.\n", os)
	}
	
	fmt.Printf("\ncount = %v", count)
}

func switch_2() {
	fmt.Println("\nWhen is Saturday?")

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Day After Tomorrow")
	default:
		fmt.Println("Not near :(")
	}
}

func switch_3(n int) {
	var ans bool

	// with no cond given in the switch case, it will act like a if-else chain
	switch {
	case n % 2 == 0:
		ans = false
	case n % 2 != 0:
		ans = true
	default:
		ans = false
	}

	fmt.Println(ans)
}

func main() {
	switch_1()
	switch_2()
	switch_3(5)
}