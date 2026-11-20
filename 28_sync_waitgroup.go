//sync.WaitGroup is used to wait for a group of goroutines to finish executing before the main goroutine exits.

// why-Without a WaitGroup, the main function may finish before the goroutines complete their work.

// Add(n)-Adds the number of goroutines to wait for (wg.Add(2))
// Done()-Called when a goroutine finishes. (defer wg.Done())It decreases the internal counter by 1.
// Wait()-Blocks until the counter becomes 0.(wg.Wait())

// What happens if you forget Done()?
// The counter never reaches zero, so Wait() blocks forever, causing a deadlock.

// A common pattern is defer wg.Done() at the beginning of the goroutine so it's always executed, even if the function returns early.

package main

import (
	"fmt"
	"sync"
)

func printMessage(wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this goroutine is finished

	fmt.Println("Hello from Goroutine")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // We are going to start 1 goroutine

	go printMessage(&wg)

	wg.Wait() // Wait until the goroutine calls Done()

	fmt.Println("Main function completed")
}
