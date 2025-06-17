package main

import (
	"fmt"
	"pkg4/mathutil"
)

func main() {
	a, b := 5, 25
	fmt.Printf("Sum of %d and %d is: %d\n", a, b, mathutil.Sum(a, b))
	fmt.Printf("Minus of %d and %d is: %d\n", a, b, mathutil.Minus(a, b))
	fmt.Printf("Diff of %d and %d is: %d\n", a, b, mathutil.Diff(a, b))
	fmt.Printf("Factorial of %d is: %d\n", a, mathutil.Factorial(a))
	fmt.Printf("Factorial of %d is: %d\n", b, mathutil.Factorial(b))
}
