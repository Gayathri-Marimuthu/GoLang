// package main

// import (
// 	"fmt"
// 	// "time"
// 	"sync"
// )

//Basic
// func printMessage() {
// 	fmt.Println("Welcome to goroutine")
// }

// func main() {

// 	go printMessage()

// 	time.Sleep(time.Second)

// 	fmt.Println("main function")
// }

// Multiple goroutine
// func worker(id int) {
// 	fmt.Println("Worked Id:", id)
// }

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		go worker(i)

// 	}
// 	time.Sleep(time.Second)
// }

//with waitgroup

package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker:", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All Workers Finished")
}
