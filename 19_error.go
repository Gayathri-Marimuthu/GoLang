// An error is a value that indicates something went wrong during program execution. Go represents errors using the built-in error interface.

// errors.New()-It creates a new error with a message.

// For example:

// File not found
// Database connection failed
// Invalid input
// Division by zero (if you choose to represent it as an error)

// syntax : result, err := function()

package main

import (
	"errors"
	"fmt"
)

func login(password string) error {
	if password != "gayathri123" {
		return errors.New("Invalid Password...!")
	}
	return nil
}

func main() {
	err := login("gayathri")
	if err != nil {
		fmt.Println("Login Sucess")
	}
}
