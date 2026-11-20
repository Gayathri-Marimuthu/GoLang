// Recursion is a technique where a function calls itself to solve a problem.

// Instead of using a loop (for), the function keeps calling itself until a base condition is met.

package main

import "fmt"

// formula: n *(n-1)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}
