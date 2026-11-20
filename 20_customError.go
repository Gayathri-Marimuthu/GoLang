// A custom error is an error type that you create yourself. It's useful when you want to include more information than just a plain error message.

package main

import "fmt"

type AgeError struct {
	Age int
}

func (e AgeError) Error() string {
	return fmt.Sprintf("Age %d is not eigible", e.Age)
}

func checkAge(age int) error {
	if age < 18 {
		return AgeError{Age: age}
	}
	return nil
}

func main() {
	err := checkAge(16)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Eligible")
}

Question

Interview answer

What is error handling in Go?						Explicitly checking and handling errors returned by functions.
Does Go use try-catch for ordinary error handling?	No. Go typically returns errors as values.
What does err != nil mean?							An error value was returned and needs to be handled.
What does errors.New() do?							Creates a simple error value.
What does fmt.Errorf() do?							Creates a formatted error; %w can wrap another error.
Should we ignore errors?							Generally no. Handle them or deliberately document why they can be ignored.