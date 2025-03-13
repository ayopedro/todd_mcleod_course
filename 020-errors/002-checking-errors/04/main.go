package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("names.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	bs, err := io.ReadAll(f)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
