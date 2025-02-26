package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	// a struct is a composite data type - complex data types
	// struct allows us to compose together values of different types

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	p2 := person{
		firstName: "Jenny",
		lastName:  "Moneypenny",
		age:       27,
	}

	fmt.Println(p1)
	fmt.Printf("%T\t %#v\n", p1, p1)
	fmt.Println(p2)
	fmt.Printf("%T\t %#v\n", p2, p2)

	fmt.Println(p1.firstName, p1.lastName, p1.age)

	// embedded struct
	// embedding one struct into another struct
	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}

	fmt.Println(sa1)

}
