package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 5 =", mySum(4, 5))
	fmt.Println("5 + 7 =", mySum(5, 7))
}

func mySum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}

	return sum
}

/*
	- doing a series of tests.
	- it allows us to test a variety of situations
*/
