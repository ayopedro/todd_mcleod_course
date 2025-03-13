package main

import (
	"fmt"

	"github.com/ayopedro/todd_mcleod_course/021-testing-benchmarking/05-coverage/sayings"
)

func main() {
	fmt.Println(sayings.Greet("Ayotunde"))
}

/*
	basically refers to how much of the code is covered by test
	- go test -cover

	- benchmark: 70-80%
	- to save the output to a file: go test -coverprofile cover.out
	- to open in a browser: go tool cover -html=cover.out
*/
