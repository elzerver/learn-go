package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), 10)

	if err != nil {
		panic(err)
	}
	fmt.Println(bs)

	// Comparing hashed password
	badPassword := `123pepe`

	err = bcrypt.CompareHashAndPassword(bs, []byte(badPassword))
	if err != nil {
		fmt.Println("You can't login\n", err)
	}
	fmt.Println("It seems that is correct")

}
