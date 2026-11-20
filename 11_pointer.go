// A pointer is a variable that stores the memory address of another variable instead of storing the actual value.
// -we use the pointer when we need to modify the original value,avoid copy the large structs
// -& operator gets variable address * access and modify the value
// -go does not support pointer arithmetic, making it safer than languages like C
// -Although Go always passes arguments by value, passing a pointer lets a function modify the original data because the copied pointer still points to the same memory location.
// Declaration var p *int

package main

import "fmt"

func main() {
	x := 10

	p := &x // Pointer stores the address of x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Pointer p:", p)
	fmt.Println("Value through pointer:", *p)

	// x := 10

	// p := &x

	// *p = 20 // Change the value through the pointer

	// fmt.Println("x:", x)
}

So don't think:

"Pointer is used whenever I want to change a value." ❌

Think:

"Pointer is useful when I need to give someone/function access to the original value." ✅
