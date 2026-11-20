// Ticker -A ticker repeatedly triggers an event at a fixed time interval until it is stopped.

package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for i := 1; i <= 3; i++ {
		<-ticker.C
		fmt.Println("Tick", i)
	}
}

// o/p:
// (after 2 sec)
// Tick 1

// (after 2 sec)
// Tick 2

// (after 2 sec)
// Tick 3

// ticker.C is a channel.

package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Running every second")
	}
}

// o/p:Running every second
// Running every second
// Running every second
// ...