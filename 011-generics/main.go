package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// redundant code
func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// generic func
func addG[T int | float64](a, b T) T {
	return a + b
}

// type set
type myType interface {
	~int | ~float64
}

func addG2[T myType](a, b T) T {
	return a + b
}

// type alias
type myAlias int

// package constraints
type myTypeC interface {
	constraints.Integer | constraints.Float
}

func addG3[T myTypeC](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(3, 5))
	fmt.Println(addF(3.2, 4.4))
	fmt.Println("-----Generics----")
	fmt.Println(addG(7, 9))
	fmt.Println(addG(7.3, 9.2))
	fmt.Println(addG2(3, 4.5))

	var n myAlias = 42
	fmt.Println(addG2(n, 4))
	fmt.Println(addG3(n, 6))
}

/*
	- Generics help us to keep our code DRY

	Understanding ~ in Type Sets (Go Generics)

	The ~ (tilde) operator in Go is used in type constraints to allow underlying types in interfaces.

	Breaking It Down

	type myType interface {
		~int | ~float64
	}

	🔹 This means any type whose underlying type is int or float64 is allowed.

	Example Usage

	type Age int       // Underlying type is int
	type Score float64 // Underlying type is float64

	func printValue[T myType](val T) {
		fmt.Println(val)
	}

	func main() {
		var a Age = 25
		var s Score = 99.5

		printValue(a) // ✅ Allowed (Age's underlying type is int)
		printValue(s) // ✅ Allowed (Score's underlying type is float64)
	}

	Key Takeaways

	✔ ~T allows custom types with the same underlying type.
	✔ myType includes int, float64, and their aliases.
	✔ Without ~, only int or float64 (not their aliases) would be allowed.

*/
