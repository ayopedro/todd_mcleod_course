package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(49, 42, add)
	y := doMath(92, 31, subtract)

	fmt.Println(x)
	fmt.Println(y)
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

/*
	note:

	a callback function is a function passed as a parameter into another function
*/
