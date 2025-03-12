package main

import "fmt"

func main() {
	c := make(<-chan int, 2) // a channel that only sends i.e., receive only channel
	fmt.Printf("%T\n", c)

	c <- 42 // fails because we are trying to send to a channel that we should only receive values from
	c <- 43 // error: cannot send to receive-only channel c (variable of type <-chan int)

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
