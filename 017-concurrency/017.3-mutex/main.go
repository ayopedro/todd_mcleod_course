package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("CPUS:%d\n", runtime.NumCPU())
	fmt.Printf("Goroutines before:%d\n", runtime.NumGoroutine())

	counter := 0

	const gs = 1000
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for range gs {
		go func() {
			defer wg.Done()
			mu.Lock() // lock the code below
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // unlock the code
		}()
		fmt.Printf("Goroutines during:%d\n", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("Goroutines after:%d\n", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}

/*
	A Mutex (mutual exclusion) is used to prevent race conditions by allowing only one goroutine to access a shared resource at a time.

	A race condition occurs when **multiple goroutines access and modify shared data at the same time**, leading to unpredictable results.

	Go provides sync.Mutex to lock and unlock access to shared data.

	How Mutex Works
	•	mu.Lock() → Blocks other goroutines from accessing the resource.
	•	mu.Unlock() → Releases the lock so other goroutines can proceed.

	When to Use sync.Mutex?
	•	When multiple goroutines access and modify shared memory.
	•	When you can’t avoid shared state.
	•	When using channels isn’t ideal for your use case.
*/
