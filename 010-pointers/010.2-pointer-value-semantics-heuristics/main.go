package main

import "fmt"

type Counter struct {
	value int
}

func increment(c Counter) {
	c.value++
}

func incrementP(c *Counter) {
	c.value++
}

func main() {
	c1 := Counter{value: 10}
	c2 := Counter{value: 10}

	increment(c1)
	fmt.Println("After increment:", c1.value)

	incrementP(&c2)
	fmt.Println("After incrementP:", c2.value)

}

/*
	Guidelines
	- use value semantics when possible
		-> simpler and safer
		-> if a function does not need to modify its input or the data you are working with is small (like built-in types or small struct), use value semantics
	- use pointer semantics for large data
		-> copying large structs or arrays can be inefficient.
		-> if the data you are working with is large (64-bytes or larger), then use pointers
	- use pointer semantics for immutability
		-> if a function needs to modify its receiver or input parameter, you should use pointer semantics
		-> common use case for methods that needs to update the state of a struct
	- consistency
		-> it is important to be consistent. once a type has a method that uses pointer semantics, all methods on that type must use pointer semantics
	- Pointer semantics when interfacing with other codes.
		-> when you are interfacing with other codes (libraries or a system call), you might need to use pointer semantics
*/
