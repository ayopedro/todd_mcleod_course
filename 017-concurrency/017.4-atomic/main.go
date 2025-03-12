package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Printf("CPUS:%d\n", runtime.NumCPU())
	fmt.Printf("Goroutines before:%d\n", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
		}()
		fmt.Printf("Goroutines during:%d\n", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("Goroutines after:%d\n", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
