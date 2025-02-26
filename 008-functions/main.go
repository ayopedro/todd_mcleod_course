package main

import "fmt"

func main() {
	/*
		syntax: func (r receiver) identifier(p parameter(s))(return(s)){ <code /> }

		- receivers would give a type of method
		- everything in Go is PASS BY VALUE
		- identifier: name of the function
	*/
	foo()
	bar("Mickey")
	s := aloha("Ayotunde")

	fmt.Println(s)

	t, a := dogYears("Rooney", 3)
	fmt.Println(t, a)

}

func foo() {
	fmt.Println("I am from Foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("My name is ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years:"), age
}
