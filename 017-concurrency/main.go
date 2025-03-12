package main

func main() {}

/*
	Concurrency vs. Parallelism in Go

	🔹 Concurrency → Doing many tasks by switching between them (interleaving).
	🔹 Parallelism → Doing many tasks at the same time (on multiple cores).

	Concurrency (Task Switching)

	Go achieves concurrency using goroutines (lightweight threads).
	Example:

	package main

	import (
		"fmt"
		"time"
	)

	func task(name string) {
		for i := 1; i <= 3; i++ {
			fmt.Println(name, i)
			time.Sleep(500 * time.Millisecond) // Simulates work
		}
	}

	func main() {
		go task("Task 1")
		go task("Task 2")
		time.Sleep(2 * time.Second) // Wait for goroutines
	}

	✔ The tasks run concurrently but may not execute at the same time.

	Parallelism (Multi-core Execution)

	To achieve real parallel execution, Go uses multiple CPU cores.
	Example:

	package main

	import (
		"runtime"
		"fmt"
	)

	func main() {
		fmt.Println("Available CPUs:", runtime.NumCPU()) // Shows CPU cores
		runtime.GOMAXPROCS(runtime.NumCPU())            // Use all cores
	}

	✔ This enables true parallel execution when running multiple goroutines.

	Key Differences

	Feature	| Concurrency	| Parallelism
	Execution	| Task switching	| Running simultaneously
	CPU Usage	| Single/multiple	| Multiple cores
	Example	| Goroutines (go func())	| runtime.GOMAXPROCS(n)
	Analogy	| Chef cooking multiple dishes (switching tasks)	| Multiple chefs cooking at once

*/
