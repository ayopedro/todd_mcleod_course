package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("exiting code")
}

// send channel
func send(e, o chan<- int) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

// receive channel
func receive(e, o <-chan int, f chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range e {
			f <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range o {
			f <- v
		}
	}()

	wg.Wait()
	close(f)
}
