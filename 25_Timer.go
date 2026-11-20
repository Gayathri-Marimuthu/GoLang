// **Timer
// A timer is used to trigger an event after a specified duration. We create it using time.NewTimer(). Every timer has a channel called timer.C. When the timer expires, it sends the current time on that channel. For simple one-time timeouts, we can use time.After(). If we need to cancel a timer before it expires, we use timer.Stop().

// **Suppose you want to:
// -Retry an API after 5 seconds.
// -Log out a user after 30 minutes.
// -Delay a task for 10 seconds.

// time.NewTimer() also creates a timer, but you get a Timer object that you can stop or reset.
// time.After(d) It waits for the specified duration and then sends the current time on a channel.

// time.NewTimer()  → Create timer
// timer.C          → Wait for expiry
// timer.Stop()     → Cancel timer
// time.After()     → Simple one-time timer
// select           → Timeout handling

package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(3 * time.Second)

	fmt.Println("Waiting...")

	<-timer.C

	fmt.Println("Timer Expired")
}

// o/p:
// Waiting...

// (after 3 seconds)

// Timer Expired
