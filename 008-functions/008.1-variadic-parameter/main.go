package main

import "fmt"

func main() {
	// variadic parameter: a func that takes an unlimited number of arguments. ... operator

	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("The sum is", x)

	// unfurling a slice
	xs := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	x1 := sum(xs...)
	fmt.Println("The sum is", x1)
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
