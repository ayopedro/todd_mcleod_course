package main

import "fmt"

// value semantics
func addOne(n int) int {
	return n + 1
}

// pointer semantics
func addOneAndChange(n *int) {
	*n += 1
}

func main() {
	a := 1

	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	addOneAndChange(&a)
	fmt.Println(a)
}

/*
	Value vs. Pointer Semantics in Go

	In Go, variables can be passed and assigned using value semantics or pointer semantics.

	1. Value Semantics (Copy Behavior)
		•	When you assign or pass a value, a copy is created.
		•	Changes do not affect the original value.

	Example: Passing by Value (Immutable)

	func changeValue(x int) {
		x = 100 // Only changes local copy
	}

	func main() {
		num := 42
		changeValue(num)
		fmt.Println(num) // Still 42
	}

	Here, num is copied into x, so changes inside changeValue don’t affect the original num.

	2. Pointer Semantics (Reference Behavior)
		•	Instead of passing a copy, a memory address (pointer) is passed.
		•	Changes affect the original value.

	Example: Passing by Pointer (Mutable)

	func changeValue(x *int) {
		*x = 100 // Modifies original value
	}

	func main() {
		num := 42
		changeValue(&num) // Pass memory address
		fmt.Println(num)  // Now 100
	}

		•	&num gets the address of num.
		•	x *int is a pointer that points to num.
		•	*x = 100 updates the actual value stored at that address.

	Key Differences

	Aspect	Value Semantics	Pointer Semantics
	How it’s passed	As a copy	As a reference (pointer)
	Memory impact	More memory (new copy)	Less memory (shared reference)
	Mutation effect	Does not affect original	Affects original value
	Used in	Simple types, small structs	Large structs, shared state

	When to Use?

	✔ Use value semantics when:
		•	You don’t need modifications.
		•	You’re working with small data types (e.g., int, string).

	✔ Use pointer semantics when:
		•	You want to modify the original value.
		•	You’re dealing with large structs (to avoid unnecessary copying).

	Conclusion
		•	Value semantics → Copies data (safe but can be inefficient).
		•	Pointer semantics → Shares data (efficient but can cause unintended modifications).

	Go defaults to value semantics but allows pointer semantics when needed. 🚀
*/
