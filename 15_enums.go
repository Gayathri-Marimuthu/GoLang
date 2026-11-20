//Go does not have a built-in enum keyword like Java or C#.
//Instead, Go uses custom types and the const keyword with iota to create enum-like values.

//An enum (enumeration) is a type that can have a fixed set of predefined values.

//What is iota?
//iota is a predeclared identifier used inside a const block. It starts at 0 and increments by 1 for each new constant.

package main

import "fmt"

type Day int // day is type

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	var today Day = Wednesday

	fmt.Println(today)
}
