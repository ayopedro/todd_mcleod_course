package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send values to the channels
	go send(even, odd, quit)

	// receive values from the channels
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:\t", v)
		case v := <-o:
			fmt.Println("from the odd channel:\t", v)
		case v := <-q:
			fmt.Println("from the quit channel:\t", v)
			return
		}
	}
}
