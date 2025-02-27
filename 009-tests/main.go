package main

import "fmt"

func main() {

	a := Add(7, 12)
	dm := doMath(7, 12, Add)
	dm2 := doMath(12, 7, Subtract)

	f := factorial(5)

	fmt.Println(a)
	fmt.Println(dm, dm2)
	fmt.Println(f)
}

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// doMath return a result based on the function passed into it
func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

// Subtract returns the difference between two numbers
func Subtract(a, b int) int {
	return a - b
}

// factorial returns the multiplication of all numbers leading to the specified argument.
func factorial(v int) int {
	if v <= 1 {
		return 1
	}

	return v * factorial(v-1)
}
