package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := range 100 {
			c <- i
		}

		close(c) // code breaks if the channel is not closed
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit...")
}

/*
	Big pieces
	- channels block
	- range also blocks
*/
