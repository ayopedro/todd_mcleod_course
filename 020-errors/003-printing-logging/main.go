package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Open("no-file.txt")

	if err != nil {
		// fmt.Println("error occured:", err)
		// log.Println("error occured:", err)
		log.Fatalln(err) // deferred functions are not run. the program exits
		// panic(err)
	}
}
