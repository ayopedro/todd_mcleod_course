package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and i am walking")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and i am running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func youngWalk(y youngin) {
	y.walk()
}

func main() {
	d1 := dog{
		first: "Jack",
	}

	d1.walk()
	d1.run()
	// youngRun(d1)
	// youngWalk(d1)

	d2 := &dog{"Henry"}
	d2.run()
	youngRun(d2)
	youngWalk(d2)
}

/*
	- a method set is a set of methods attached to a type.
	- this concept is key to Go's interface mechanism
	- it is associated to both pointer and value types.
	- the method set of a type T consists of all methods with receiver type T
		-> these methods can be called using variables of type T
	- the method set of a type *T consists of all methods with receiver *T or T
		-> these methods can be called using variables of type *T
		-> can also be called using corresponding non-pointer types as well
*/
