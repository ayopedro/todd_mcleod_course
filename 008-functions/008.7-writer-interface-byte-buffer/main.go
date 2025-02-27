package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	// _, err := w.Write([]byte(p.first))

	// return err
	w.Write([]byte(p.first))
}

func main() {
	p := person{
		first: "Jimmy",
	}

	f, err := os.Create("output.txt")

	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)

	fmt.Println(b.String())

	// s := []byte("Hello Gophers!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	// b := bytes.NewBufferString("Hello ")

	// fmt.Println(b.String())

	// b.WriteString("Gophers!")
	// fmt.Println(b.String())

	// b.Reset()
	// b.WriteString("It's Wednesday")
	// fmt.Println(b.String())

	// b.Write([]byte("Happy happy"))
	// fmt.Println(b.String())
}

/*
	notes:

	- a byte buffer is a region of memory used to temporarily store a sequence of bytes.
	- it provides and efficient manipulation of byte data.
	- a byte buffer allows you to read and write bytes to and from the buffer making
	it useful for tasks like data serialization, network communication, file I/O and
	efficient string manipulation
	- it is not specific to any programming language.
	- a byte buffer is a data structure that provides a convenient interface to manipulate
	sequences of bytes efficiently. it serves as a temporary storage for byte data and enables
	operations such as reading, writing, appending and resizing byte structures.
*/
