package test2

// Reverse reverses the digits of an integer
func Reverse(n int) int {
    reversed := 0
    for n > 0 {
        remainder := n % 10
        reversed = reversed*10 + remainder
        n /= 10
    }
    return reversed
}