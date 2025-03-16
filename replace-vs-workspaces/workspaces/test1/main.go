package main

import (
    "fmt"
    "halogen/test2"
)

func main() {
    num := 12345
    reversed := test2.Reverse(num)
    fmt.Printf("Original: %d, Reversed: %d\n", num, reversed)
}