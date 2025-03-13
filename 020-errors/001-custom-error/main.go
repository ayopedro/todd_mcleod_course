package main

import "fmt"

type CustomError struct {
	msg string
}

func (ce *CustomError) Error() string {
	return ce.msg
}

func main() {

	err := &CustomError{"Invalid credentials"}

	fmt.Printf("%s\n", err)
}
