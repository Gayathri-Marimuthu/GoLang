A channel is a way for goroutines to communicate safely.

Instead of sharing variables directly, one goroutine sends data through a channel, and another goroutine receives it.

A channel is a typed communication mechanism in Go that allows goroutines to safely send and receive data. It also provides synchronization between goroutines

<!-- Without a channel: -->

Goroutine 1 -----> Shared Variable <----- Goroutine 2

Both goroutines can access the same variable at the same time.
This can cause a data race.

<!-- With a channel: -->

Goroutine 1 -----> Channel -----> Goroutine 2

<!-- creating channel -->

syntax: ch:=make(chan int) --This creates a channel that carries int values.
strCh := make(chan string)
boolCh := make(chan bool)

<!-- sending data-->

Use <-.

ch <- 100 --This sends 100 into the channel.

<!-- receiving data -->

value := <-ch

fmt.Println(value)

<!-- close channel -->

close(ch)

<!-- reading from closed channel -->

value, ok := <-ch

if closed value=0 ,ok=false

<!-- Real example of Race condition -->

Imagine two people sharing a bank account with $100.

Person A withdraws $20.
Person B withdraws $30.

If both ATMs read the balance ($100) simultaneously:

ATM A computes new balance = $80.
ATM B computes new balance = $70.

If ATM A updates first and ATM B updates second, the final balance becomes $70 instead of the correct $50.

<!-- Buffered and unbuffered channel -->

\*\*\*Unbuffered channel - An unbuffered channel has no storage. A send operation blocks until another goroutine receives the value.

there is no receiver then grt below error

<!-- fatal error: all goroutines are asleep - deadlock! -->

<!-- package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
		fmt.Println("Value sent")
	}()

	value := <-ch

	fmt.Println("Received:", value)
} -->

\*\*\*buffered-A buffered channel has storage.The sender can send values until the buffer becomes full.

<!-- package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
} -->No goroutine is needed because the buffer stores the values.

<!-- mind hint -->

Unbuffered = Hand-to-hand delivery. The sender must wait until the receiver takes the value.
Buffered = Mailbox. The sender can drop messages into the mailbox until it's full, and the receiver can collect them later.

<!-- Capacity -->

fmt.Println(cap(ch))
fmt.Println(len(ch))

<!-- Directional Channels -->

func send(ch chan<- int) //send only
func receive(ch <-chan int) //received only

<!-- using range -->

for value := range ch {

    fmt.Println(value)

}

\*\*\*select
select allows a goroutine to wait on multiple channel send and receive operations. When one of the channels becomes ready, Go executes the corresponding case.

<!-- default -->

No channel is ready.So default executes immediately.

<!-- Blocking and non-blocking channels -->

Normally, channel send and receive operations block until they can complete. A non-blocking channel operation allows the program to continue immediately if the channel isn't ready. In Go, we implement this using select with a default case. If a send or receive can proceed, that case executes; otherwise, the default case runs without blocking the goroutine.
