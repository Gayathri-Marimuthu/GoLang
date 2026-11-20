// Closure=Function + Remembered Variable
//A closure is a function that captures and remembers variables from its surrounding scope, even after the outer function has returned.

package main

import "fmt"

func count() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	nxtint := count()
	fmt.Println(nxtint())
}

//The outer count() function finishes, but the returned function still carries count in its bag.
