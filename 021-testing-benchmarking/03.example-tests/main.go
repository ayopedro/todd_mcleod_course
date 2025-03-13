package main

import (
	"fmt"

	"github.com/ayopedro/todd_mcleod_course/021-testing-benchmarking/03.example-tests/acdc"
)

func main() {
	s := acdc.Sum(2, 3, 4, 5, 6, 7)

	fmt.Println(s)
}
