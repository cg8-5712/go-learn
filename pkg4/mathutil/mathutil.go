package mathutil

// Compute the sum
func Sum(a, b int) int {
	return a + b
}

// Compute the minus
func Minus(a, b int) int {
	return a - b
}

// Compute the difference
func Diff(a, b int) int {
	return a / b
}

// Factorial
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
