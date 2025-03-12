package main

import "fmt"

func main() {
	c := make(chan int, 1) // a buffer channel

	c <- 42
	c <- 43 // there's no space for this to be passed on to the channel

	fmt.Println(<-c)
}

// NOTE TO SELF: try to stay away from buffer channels
