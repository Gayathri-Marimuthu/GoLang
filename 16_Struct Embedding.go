//Struct embedding means including one struct inside another without giving it a field name. The embedded struct's fields and methods are promoted, so they can be accessed directly from the outer struct.

// Struct embedding is a feature in Go that allows one struct to include another struct without using inheritance.

// It helps you reuse fields and methods from another struct.

package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Age     int
	Address // Embedded struct
}

func main() {
	emp := Employee{
		Name: "John",
		Age:  25,
		Address: Address{
			City:  "Chennai",
			State: "Tamil Nadu",
		},
	}

	fmt.Println(emp.Name) // avoid emp.Address.City
	fmt.Println(emp.City)
	fmt.Println(emp.State)
}
