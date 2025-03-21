package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	loginPW := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPW))

	if err != nil {
		fmt.Println("You can't login!")
		return
	}
	fmt.Println("You're authenticated!")
}
