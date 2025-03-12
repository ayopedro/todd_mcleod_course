package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context error:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("---------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context error:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("---------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context error:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("---------")

}

/*
	context Package in Go

	The context package is used for managing timeouts, deadlines, and cancellations in Go programs, especially for goroutines.

	⸻

	Creating a Context

	1️⃣ context.Background() (Base context)

	ctx := context.Background() // Root context

	Used as a starting point for other contexts.

	2️⃣ context.WithCancel() (Cancel on demand)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done() // Wait for cancellation
		fmt.Println("Goroutine stopped")
	}()

	cancel() // Cancels the context

	✅ Use case: Stop a goroutine when no longer needed.

	⸻

	Context With Timeout

	3️⃣ context.WithTimeout() (Auto-cancel after a time)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout reached:", ctx.Err()) // Prints "context deadline exceeded"
	}

	✅ Use case: Prevent long-running operations.

	⸻

	Context With Deadline

	4️⃣ context.WithDeadline() (Set a specific deadline)

	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	✅ Use case: Ensure an operation stops at a fixed time.

	⸻

	Passing Context to Functions

	func doWork(ctx context.Context) {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Work done")
		case <-ctx.Done():
			fmt.Println("Cancelled:", ctx.Err()) // Stops if cancelled
		}
	}

	func main() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		go doWork(ctx)
		time.Sleep(3 * time.Second) // Exceeds timeout
	}

	✅ Useful for HTTP requests, database queries, and background jobs.

*/

/*
	practice:
	- https://go.dev/play/p/y2wTQxlByAA
	- select: https://go.dev/play/p/Sdg-pYQvDFH
	- comma ok idiom: https://go.dev/play/p/DBWZ5kOzyJm
	- https://go.dev/play/p/50LPa3p40Xc

	- https://go.dev/play/p/bW7y5p3RYaN
*/
