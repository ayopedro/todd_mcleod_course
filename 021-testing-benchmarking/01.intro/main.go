package main

import "fmt"

func main() {
	xn := []float64{1, 2, 3, 4, 5}

	avg := Average(xn)

	fmt.Println("The average is:", avg)
}

func Average(n []float64) float64 {
	var sum float64
	length := len(n)
	for _, v := range n {
		sum += v
	}

	return sum / float64(length)
}

/*
	Tests must:
		- be in a file that ends with "_test.go"
		- put the file in the same package as the one being tested
		- be a func with a signature func TestXxx(*testing T)

	To run a test:
		- go test

	To deal with test failure:
		- use the Error, Fail or related methods to signal failure
*/
