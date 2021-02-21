package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	icecream string
}


func main() {
	p1 := person{
		first: "Peter",
		last: "Griffin",
		icecream: "Vanilla",
	}

	fmt.Println(p1)
}