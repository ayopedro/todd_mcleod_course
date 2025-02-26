package main

import "fmt"

type person struct {
	firstName string
}

func (p person) speak() {
	fmt.Println("I am", p.firstName)
}

func main() {
	// a method is a Func attached to a TYPE. you attach a func to a type with a RECEIVER
	p1 := person{
		firstName: "Ayotunde",
	}

	p2 := person{
		firstName: "Jenny",
	}

	p1.speak()
	p2.speak()
}
