package main

import "fmt"

func main() {
	f := increment()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
