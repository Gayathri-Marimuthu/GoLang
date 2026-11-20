// A struct is a user-defined data type in Go that groups related fields of different data types into a single unit. It is used to model real-world objects.

package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	emp := Employee{
		Name:   "John",
		Age:    25,
		Salary: 50000,
	}

	fmt.Println(emp)
}

//Accessing Struct Fields
// fmt.Println(emp.Name)
// 1. Named Fields (Recommended)
// emp := Employee{
// 	Name:   "John",
// 	Age:    25,
// 	Salary: 50000,
// }

// 2. Positional Fields
// emp := Employee{"John", 25, 50000}

// 3.Zero value
// var emp Employee
// fmt.Println(emp)

// 4.Updating Struct Fields
// emp.Age = 30
// fmt.Println(emp)

// 5.Anonymous Struct
// person := struct {
// 	Name string
// 	Age  int
// }{
// 	Name: "John",
// 	Age:  25,
// }
// fmt.Println(person)

// 6.Nested Struct
