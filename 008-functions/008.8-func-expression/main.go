package main

import "fmt"

func main() {
	v := foo()
	b := bar()

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}

	fmt.Println(v)
	fmt.Println(b())

	x()
	y("Ayotunde")

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", b)
}

// returning a value
func foo() int {
	return 42
}

// returning a function
func bar() func() int {
	return func() int {
		return 43
	}
}

/*
	notes

	- an expression is a combination of values, variables, operators and function calls that are
	evaluated to produce a single value
	- it can be as simple as a literal value or a variable or it can involve complex operations
	and function invocations
*/
