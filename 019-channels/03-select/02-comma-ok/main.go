package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)
}

func send(e, o, q chan<- int) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:\t", v)
		case v := <-o:
			fmt.Println("from the odd channel:\t", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("from comma ok:\t\t", v, ok)
			} else {
				fmt.Println("from comma ok:\t\t", v)
			}
			return
		}
	}
}
