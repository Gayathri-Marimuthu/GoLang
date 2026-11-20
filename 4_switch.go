package main

import "fmt"

// he switch statement is used to execute different blocks of code based on the value of an expression. It is often a cleaner and more readable alternative to multiple if...else if statements.
func main() {
	// 1.Basic switch syntax
	switch expression {
	case value1:
		// code
	case value2:
		// code
	default:
		// code
	}
	// 2.switch without expression
	// 3.multiplevalue in one case
	// 4.short variable declaration

	// 5.fallthrough
	// Normally, Go stops after executing one matching case.
	// Using fallthrough forces execution of the next case.

	// case 1:
	//     fmt.Println("One")
	//     fallthrough
}
