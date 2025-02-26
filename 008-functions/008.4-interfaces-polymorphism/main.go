package main

import "fmt"

type person struct {
	firstName string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.firstName)
}

func (sa secretAgent) speak() {
	fmt.Println("I am secret agent", sa.firstName)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	/*
		- An interface in Go refers to a set of method signatures
		- Polymorphism is the ability of a VALUE of a certain TYPE to also be of another TYPE.
		- Polymorphism refers to the ability of an object to be of an additional TYPE and take on different behaviors.
		- in the context of programming, it allows VALUES of different TYPES to be treated as instances of a common TYPE

		in Go, values can be of more than one type.
	*/

	p1 := person{
		firstName: "Ayotunde",
	}

	p2 := person{
		firstName: "Jenny",
	}

	sa1 := secretAgent{
		person: person{
			firstName: "Bond",
		},
		ltk: true,
	}

	// p1.speak()
	// p2.speak()

	// sa1.speak()

	saySomething(p1)
	saySomething(p2)
	saySomething(sa1)
}
