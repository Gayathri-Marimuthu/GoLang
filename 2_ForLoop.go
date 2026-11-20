package main

import "fmt"
// The for loop is the only looping construct in Go. It can be used like a for, while, and infinite loop.

// while loop:a control structure that repeatedly executes a block of code as long as a specified condition remains true

func main(){
	for i:=1;i<=5;i++{
		fmt.Println("Increament value",i)
	}

	// while loop 
	j:=1
	for j<=5{
		fmt.Println("while loop",j)
		j++
	}

	//infinity loop
	for {
		fmt.Println("infinity hello")
	}

	// Nested loop

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("Nested Loop",i, j)
		}
	}

	// continue,break,range

}
// break exits the loop.
// continue skips the current iteration and continues with the next.

