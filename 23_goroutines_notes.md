// A goroutine is a lightweight thread managed by the Go runtime. It allows multiple functions to run concurrently. We create a goroutine by using the go keyword before a function call. Unlike operating system threads, goroutines are very lightweight because they start with a small stack and are scheduled by the Go runtime instead of the OS. This allows Go applications to efficiently handle thousands of concurrent tasks.

<!-- Why Do We Need Goroutines? -->

// Without goroutines, functions execute one after another, so if one task takes a long time, the whole program waits. Goroutines allow multiple independent tasks to run concurrently, improving responsiveness and making better use of CPU resources.

// Concurrency means dealing with multiple tasks at the same time, while parallelism means executing multiple tasks simultaneously. Goroutines provide concurrency, and if multiple CPU cores are available, the Go scheduler can also run goroutines in parallel.
eg: Concurrency (One Person) ,Parallelism (Three People)

<!-- | Goroutine | Thread |
| ------------------------------ | ----------------------- |
| Managed by Go runtime | Managed by OS |
| Very lightweight | Heavyweight |
| Starts with ~2 KB stack | Usually 1–2 MB stack |
| Thousands or millions possible | Limited by OS resources |
| Faster to create | Slower to create | -->

 <!-- Why shouldnt you use time.Sleep() to wait for Goroutines? -->

-time.Sleep() only delays execution and doesnt guarantee that goroutines have finished. sync.WaitGroup is the correct synchronization mechanism.

 <!-- What is a Goroutine leak? -->

-A goroutine leak occurs when a goroutine never exits (for example, its blocked forever on a channel or waiting indefinitely), causing unnecessary memory and resource usage.

A goroutine leak happens when a goroutine is created but never finishes. It stays in memory forever, waiting for something that will never happen. Over time, many leaked goroutines can consume memory and slow down the application.

<!-- Common Causes -->

Waiting forever on a channel.
Deadlocks.
Infinite loops without an exit condition.
Network or database calls that never return.
Forgetting to cancel a goroutine using a context.

How to Prevent Goroutine Leaks
Use context.Context to cancel goroutines.
Close channels when no more data will be sent.
Use select with a timeout or cancellation.
Make sure every goroutine has a clear exit condition.

<!-- | Process | Thread | Goroutine |
| ------------------- | ------------------------------- | ---------------------------------------- |
| Independent program | Execution unit within a process | Lightweight execution unit managed by Go |
| Has its own memory | Shares process memory | Shares process memory |
| Managed by OS | Managed by OS | Managed by Go runtime | -->

 <!-- Go Runtime does something like this (conceptually): -->

// CPU

// ↓

// Go Scheduler -The Go Scheduler decides which goroutine gets CPU time. eg teacher decided who is speak first

// ↓

// -------------------------
// | Main Goroutine |
// -------------------------

// -------------------------
// | task Goroutine |
// -------------------------

 <!-- When Should You Use Goroutines? -->

Use goroutines when tasks are independent, for example:
Calling multiple APIs at the same time.
Processing multiple files.
Running background jobs.
Reading from Kafka while handling HTTP requests.
Sending emails without blocking the main request.
Running database operations concurrently when appropriate.

<!-- How Does a Goroutine Work Internally? -->

When we use the go keyword, Go creates a new goroutine. The Go runtime scheduler is responsible for deciding when each goroutine runs. Multiple goroutines are multiplexed(Many goroutines share a small number of OS threads.) onto a smaller number of operating system threads, which makes them much lighter than creating an OS thread for every task.

<!-- Why Are Goroutines Lightweight? -->

Goroutines are lightweight because they start with a very small stack(memory area), around 2 KB, and the stack grows and shrinks automatically as needed. Also, the Go runtime manages scheduling, so creating and switching goroutines is much cheaper than creating OS threads.

<!-- What Problems Can Occur with Goroutines? -->
<!-- -Data Race -->

Multiple goroutines access the same variable, and at least one modifies it without synchronization.
Solution: sync.Mutex, channels, or atomic operations.

<!-- -Goroutine Leak -->

A goroutine never exits because its blocked forever (for example, waiting on a channel that never receives data).

<!-- -Improper Synchronization -->

If main() exits before goroutines finish, they are terminated.
Solution: sync.WaitGroup or channels.

<!-- Real time example -->

Suppose an e-commerce application receives an order. After saving the order to the database, it also needs to send a confirmation email, generate an invoice, and update analytics. These tasks are independent, so instead of doing them one by one, we can run them in separate goroutines. This improves response time because the user doesnt have to wait for every background task to complete.

<!-- Why do we use WaitGroup? -->

WaitGroup synchronizes goroutines by making the main goroutine wait until all worker goroutines have finished.

<!-- Why do we use the go keyword? -->

The go keyword starts a function in a new goroutine, allowing it to execute concurrently with the current goroutine.
