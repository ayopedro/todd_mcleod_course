package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("this is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 008:", s.String())
}

func main() {
	b := book{
		title: "West with the Night",
	}

	var c count = 62

	logInfo(b)
	logInfo(c)
}

/*
	Notes:

	- a wrapper function is a function that provides an additional layer of abstraction or functionality around an existing
	function or method
	- it acts as an intermediary between the caller and the wrapped function allowing you to modify inputs, outputs or behavior
	without directly modifying the original function
*/
