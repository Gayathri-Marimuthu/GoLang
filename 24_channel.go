package main

import "fmt"

func send(ch chan int) {
	ch <- 100
}

func main() {
	ch := make(chan int)
	go send(ch)
	value := <-ch
	fmt.Println(value)
}

// close channel
func main() {

	ch := make(chan int)

	go func() {

		ch <- 1
		ch <- 2

		close(ch)

	}()

	for value := range ch {

		fmt.Println(value)
	}
}
// channel with waitgroup
package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(2 * time.Second)
	fmt.Println("Hello")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go printHello(&wg)

	wg.Wait()

	fmt.Println("Main Finished")
} //without waitgrop main execute first

// Select
package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg := <-ch2:
		fmt.Println(msg)
	}
// case <-time.NewTimer(2 * time.Second).C:
// 	fmt.Println("Timeout")

}//o/p:Hello from ch2
// Note:Although ch1 also sends data later, select has already completed because it executes only one matching case.