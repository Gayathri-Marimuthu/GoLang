package main

import "fmt"

// The if statement is used to make decisions in Go. Based on whether a condition is true or false, Go executes different blocks of code.
func main() {
	// 1.simple if
	age := 20
	if age >= 18 {
		fmt.Println("Eligble to vote")
	}
	//2. if..else
	if age >= 15 {
		fmt.Println("you are child")
	} else {
		fmt.Println("Teen Age")
	}
	//3.if..else if..else
	//4.Multiple condition || &&
	//5.short variable declar
	if num := 10; num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("odd Number")
	}
	// 6.Nested if
	salary := 50000
	if age >= 18 {
		if salary >= 30000 {
			fmt.Println("Loan Approved")
		} else {
			fmt.Println("Salary too low")
		}
	}
}
