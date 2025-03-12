package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GORoutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GORoutines\t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := range 10 {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := range 10 {
		fmt.Println("Bar:", i)
	}
}

/*
	WaitGroups in Go (sync.WaitGroup)

	🔹 sync.WaitGroup is used to wait for multiple goroutines to finish before the program exits.
	🔹 It prevents the main function from terminating before goroutines complete execution.

	How It Works
		1.	Add(n) → Increments the counter (n = number of goroutines).
		2.	Done() → Decrements the counter when a goroutine completes.
		3.	Wait() → Blocks execution until the counter reaches zero.

	Example: Using sync.WaitGroup

	package main

	import (
		"fmt"
		"sync"
		"time"
	)

	func worker(id int, wg *sync.WaitGroup) {
		defer wg.Done() // Mark as done when finished
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d done\n", id)
	}

	func main() {
		var wg sync.WaitGroup

		for i := 1; i <= 3; i++ {
			wg.Add(1)        // Increment counter
			go worker(i, &wg) // Start goroutine
		}

		wg.Wait() // Wait for all workers to finish
		fmt.Println("All workers completed!")
	}

	Output

	Worker 1 starting
	Worker 2 starting
	Worker 3 starting
	Worker 1 done
	Worker 2 done
	Worker 3 done
	All workers completed!

	Key Takeaways

	✅ Prevents main from exiting too early
	✅ Ensures all goroutines finish before continuing
	✅ Must pass a pointer (*sync.WaitGroup) to goroutines

*/
