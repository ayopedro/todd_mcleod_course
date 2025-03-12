package main

import "fmt"

func main() {
	c := make(chan int, 2) // bi-directional i.e., send and receive

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}

/*
	NB: read it from left to right

	cs := make(chan<- int) // send only
	cr := make(<-chan int) // receive only

	arrow before: extract from the channel
	arrow after: send to the channel
*/
