package main

import "fmt"

func main() {
	c := make(chan int, 1) // a buffer channel

	c <- 42

	fmt.Println(<-c)
}

/*
	Buffered Channels in Go

	A buffered channel allows sending multiple values without blocking until the buffer is full.

	⸻

	Declaring a Buffered Channel

	ch := make(chan int, 3) // Channel can hold 3 values

		•	3 is the buffer size → It can store 3 values before blocking.

	⸻

	Example: Sending & Receiving from a Buffered Channel

	package main

	import "fmt"

	func main() {
		ch := make(chan int, 2) // Buffer size of 2

		ch <- 1 // Doesn't block (buffer has space)
		ch <- 2 // Doesn't block

		fmt.Println(<-ch) // Reads 1
		fmt.Println(<-ch) // Reads 2
	}

	✅ No deadlock, as we read values from the channel.

	⸻

	Blocking Behavior

	If the buffer is full, sending blocks until a value is read.
	If the buffer is empty, receiving blocks until a value is available.

	ch := make(chan int, 1)
	ch <- 10 // Allowed (buffer has space)
	ch <- 20 // Blocks (buffer full)



	⸻

	Buffered vs Unbuffered Channels

	Type	Behavior
	Unbuffered (make(chan int))	Sender blocks until the receiver reads.
	Buffered (make(chan int, N))	Sender doesn’t block until the buffer is full.



	⸻

	When to Use Buffered Channels?

	✅ Reduce blocking when sending multiple values.
	✅ Handle bursts of data efficiently.
	✅ Improve throughput by decoupling senders & receivers.

*/
