package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter++
}

func main() {
	var wg sync.WaitGroup
	fmt.Printf("Goroutines:%d\n", runtime.NumGoroutine())

	for range 1000 {
		wg.Add(1)
		go increment(&wg)
		fmt.Printf("Goroutines:%d\n", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}

/*
	A race condition happens when multiple goroutines access shared memory without proper synchronization, leading to unexpected behavior.

	❌ Issue: Multiple goroutines try to update counter at the same time, causing inconsistent results.

	Detecting Race Conditions

	Use the Go race detector:
	go run -race main.go

	It helps find unsafe concurrent access to shared memory.
*/
