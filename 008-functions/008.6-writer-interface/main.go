package main

import (
	"log"
	"os"
)

// type person struct {
// 	first string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))

// 	return err
// }

func main() {
	f, err := os.Create("output.txt")

	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	s := []byte("Hello Gophers!")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

/*
	Note:
	- a string and a []byte (byte slice) are two different types, but they are closely related and can
	be converted between each other
	- a string in Go represents a sequence of characters. it is an immutable type which means you cannot
	modify individual characters within a string.
	- string values are always interpreted as UTF-8 encoded unicode text

	- a []byte is a slice of byte where each element represents a single byte.
	- it is a mutable type so you can modify individual bytes within a byte slice.
	- it can be used to represent binary data or text in various encodings.


	CONVERSION
	- string to []byte
	s := "Hello World"
	b := []byte(s)

	- []byte to string
	by := []byte{72, 101, 108, 108, 111}
	str := string(br)
*/
