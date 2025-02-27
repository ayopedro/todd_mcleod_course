package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")

	if err != nil {
		log.Fatalf("readFile in main: %s", err)
	}

	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fn string) ([]byte, error) {
	xb, err := os.ReadFile(fn)

	if err != nil {
		return nil, fmt.Errorf("error in readFile %s", err)
	}

	return xb, nil
}
