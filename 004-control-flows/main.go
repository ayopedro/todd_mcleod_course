package main

import (
	"fmt"
	"math/rand"
)

// func init() {
// 	fmt.Println("Initializing my program...")
// }

func main() {
	// fmt.Println("My code now runs...")

	// y := 1

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("The value of i is: %d \t first for loop\n", i)
	// }

	// for y < 10 {
	// 	fmt.Printf("The value of y is: %d \t second for loop\n", y)
	// 	y++
	// }

	// for {
	// 	fmt.Printf("The value of y is: %d \t third for loop\n", y)
	// 	if y > 20 {
	// 		break
	// 	}
	// 	y++
	// }

	// for i := 0; i < 20; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}

	// 	fmt.Println("counting even numbers: ", i)
	// }

	x := rand.Intn(250)

	fmt.Printf("The value of x is: %d\n", x)

	if x <= 100 {
		fmt.Println("The value of x is less than or equal to 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("The value of x is between 101 and 200")
	} else {
		fmt.Println("The value of x is greater than 200")
	}
}
