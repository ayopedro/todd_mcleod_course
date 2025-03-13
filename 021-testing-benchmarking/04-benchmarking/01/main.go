package main

import (
	"fmt"

	"github.com/ayopedro/todd_mcleod_course/021-testing-benchmarking/04-benchmarking/01/sayings"
)

func main() {
	fmt.Println(sayings.Greet("Ayotunde"))
}

/*
	- Benchmarking allows you to measure the performance of your code

	- command: go test -bench
	- go help testflag
*/
