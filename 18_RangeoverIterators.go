// callbackfunction means-a function passed into another function as an argument, which is then executed inside that outer function to complete an action

// Range over Iterators is a feature introduced in Go 1.23 that allows the for...range loop to iterate over an iterator function, not just slices, arrays, maps, strings, or channels.

// An iterator function generates values one at a time and passes them to the for...range loop using a callback function

package main

import "fmt"

// Iterator function
func Count(yield func(int) bool) {
	for i := 1; i <= 5; i++ {
		if !yield(i) { //used for stop the execution
			return
		}
	}
}

func main() {
	for n := range Count {

		if n == 3 {
			break
		}
		fmt.Println(n)

	}
}

1. Difference between for and range

for loop																range

Repeats code based on a condition or counter.					Iterates over items in a collection.

You manage the counter yourself when needed.					Go provides the index/key and value.

Useful when you need control over the loop.						Useful when you want to process each item.
