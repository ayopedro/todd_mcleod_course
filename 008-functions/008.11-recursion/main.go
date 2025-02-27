package main

import "fmt"

func main() {
	f := factorial(10)

	fmt.Println(f)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

/*
	- recursion is when a function calls itself until it hits a base case.
	- it refers to a technique of solving a problem by breaking it down
	into smaller subproblems of the same type
*/
