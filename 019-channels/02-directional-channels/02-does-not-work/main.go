package main

import "fmt"

func main() {
	c := make(chan<- int, 2) // a channel that only receives i.e., send only channel
	fmt.Printf("%T\n", c)

	c <- 42
	c <- 43

	fmt.Println(<-c) // fails because we are trying to receive from a channel that we should only send values to
	fmt.Println(<-c) // error: cannot receive from send-only channel c

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
