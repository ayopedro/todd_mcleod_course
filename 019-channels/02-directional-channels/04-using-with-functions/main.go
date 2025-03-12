package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go send(c)

	// receive
	receive(c)

	fmt.Println("Exiting...")
}

// send
func send(c chan<- int) {
	c <- 42
}

// receive
func receive(c <-chan int) {
	fmt.Println("value received:", <-c)
}

/*
	Channels are reference types
*/
