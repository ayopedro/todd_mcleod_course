package main

import "fmt"

func main() {
	/*
		defer statement invokes a function whose execution is deferred to the moment the
		surrounding function returns either because
		1. the surrounding function executed a return statement
		2. reached the end of its function body
		3. or because the corresponding goroutine is "panicking"
	*/
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
