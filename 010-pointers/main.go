package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func mapDelta(md map[string]int, s string) {
	md[s] = 26
}

func main() {
	// x := 42

	// fmt.Printf("%T and %v\n", &x, &x)

	// y := &x

	// fmt.Printf("%T and %v\n", y, y)
	// fmt.Println(*y)
	// fmt.Println(*&x)

	// *y = 43
	// fmt.Println(x)

	a := 42
	fmt.Println(a)

	intDelta(&a)
	fmt.Println(a)

	m := make(map[string]int)
	m["James"] = 25

	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])
}

/*
	- &: Pointer
	- *: Reference

	- A pointer stores a memory address and allows indirect access to a value. It uses the & operator to get an address and the * operator to dereference it.
	- & gives you the memory address and * gives you the value at the memory address (if you have the memory address)
	- pointers are reference types

	// CHATGPT
	What Are Pointers in Go?

	A pointer is a variable that stores the memory address of another variable. Instead of holding a value directly, it points to where the value is stored.

	Declaring and Using Pointers

	package main

	import "fmt"

	func main() {
		x := 42        // Normal variable
		p := &x        // Pointer to x

		fmt.Println(x)  // 42 (value of x)
		fmt.Println(p)  // Memory address of x
		fmt.Println(*p) // 42 (dereferencing: accessing value at pointer)
	}

		•	&x → Gets the memory address of x.
		•	p := &x → p now points to x.
		•	*p → Dereferences p to get the value stored at the address.

	Modifying a Value Using Pointers

	func updateValue(ptr *int) {
		*ptr = 100  // Change value at pointer
	}

	func main() {
		x := 42
		updateValue(&x)  // Pass address of x
		fmt.Println(x)   // 100 (updated)
	}

		•	The function receives *int, meaning a pointer to an int.
		•	*ptr = 100 updates x directly.

	Why Use Pointers?
		1.	Efficiency – Avoid copying large structures.
		2.	Mutability – Functions can modify original values.
	3.	Sharing Data – Multiple parts of a program can use the same value.

	Mutability in Go

	Mutability means whether a variable’s value can be changed after it’s created.

	In Go:
		•	Values passed by value (immutable copy) → Changes don’t affect the original.
		•	Values passed by reference (mutable via pointer) → Changes affect the original.

	Example: Immutable (Pass by Value)

	func changeValue(n int) {
		n = 100
	}

	func main() {
		x := 42
		changeValue(x)
		fmt.Println(x) // Still 42 (unchanged)
	}

		•	x is copied into n, so the original x remains unchanged.

	Example: Mutable (Using Pointers)

	func changeValue(n *int) {
		*n = 100
	}

	func main() {
		x := 42
		changeValue(&x)
		fmt.Println(x) // 100 (changed!)
	}

		•	x is passed by reference using &x.
		•	*n = 100 modifies the actual value at the memory address.

	Key Takeaway
		•	Without pointers → immutable copies
		•	With pointers → mutable references

	Go avoids unintended mutability but allows controlled changes with pointers! 🚀
*/
