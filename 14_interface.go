// An interface in Go defines a set of method signatures. Any type that implements those methods automatically satisfies the interface—there's no implements keyword.
// interfaces are commonly used to achieve polymorphism
// An interface defines a set of methods. Any type that has those methods can use the interface. It helps us write flexible code without depending on a specific type.
// For example, Email and SMS both have a Send method. I can create a Sender interface with Send(), and my function can work with either Email or SMS.
package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (d Dog) Sound() {
	fmt.Println("Dag says woof")
}

func (c Cat) Sound() {
	fmt.Println("Cat says meow")
}

func main() {
	var a Animal
	a = Dog{name: "Dora"}
	a.Sound()
	a = Cat{name: "jessy"}
	a.Sound()
}
