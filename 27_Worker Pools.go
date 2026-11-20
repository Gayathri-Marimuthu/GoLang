// Worker Pool in Go is a concurrency pattern where a fixed number of goroutines (workers) process jobs from a shared queue (channel). Instead of creating one goroutine per task, you reuse a limited number of workers, which improves performance and prevents excessive memory/CPU usage.

// Why do we use a Worker Pool?

// Without a worker pool:

// If you have 100,000 tasks, creating 100,000 goroutines can consume a lot of memory.
// Too many goroutines can increase scheduling overhead.

// With a worker pool:

// Create a fixed number of workers (e.g., 10).
// Workers continuously read jobs from a channel.
// Each worker processes one job at a time.

// Advantages
// Limits the number of goroutines.
// Better CPU utilization.
// Lower memory usage.
// Prevents overwhelming databases or external APIs.
// Improves throughput for many workloads.

package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d is processing Job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start 2 workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	// No more jobs
	close(jobs)

	// Wait for workers to finish
	wg.Wait()

	fmt.Println("All jobs completed")
}

// close-This loop ends only when the channel is closed. Without closing it, workers would block forever waiting for more jobs.

// WaitGroup-To wait until all workers complete their work before the main goroutine exits.

// How would you cancel all workers if one task fails?
// Use context.Context to signal cancellation to all workers. Each worker checks ctx.Done() and exits gracefully when cancellation is requested.
