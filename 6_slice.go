// A slice is a dynamic, flexible view of an array. Unlike arrays, slices can grow or shrink in size.
// syntax var sli[]int
// length-number of element currently presented
// capacity-Maximum element before go allocate a new arrays
package main

import "fmt"

func main() {
	numbers := []int{1, 3, 5, 6, 8}
	// length&capacity
	fmt.Println("length", len(numbers))
	fmt.Println("capacity", cap(numbers))
	//append
	numbers = append(numbers, 9)
	fmt.Println("slice", numbers)
	//make
	group := make([]int, 3, 5)
	fmt.Println("slicegroup", group)
	// group = append(group, 5, 6, 7, 8)  //output [0 0 0 5 6 7 8]
	group[0] = 20
	group[1] = 30
	group[2] = 40
	fmt.Println(group)
	//copy
	c := copy(numbers, group) //c contain copied element count
	fmt.Println("copy", c, numbers, group)
}

// we use make() with a slice when we want to create a slice with a specific length and optionally a specific capacity.
