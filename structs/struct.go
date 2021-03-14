package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "Peter",
		last:  "Parker",
		age:   27,
	}

	fmt.Println("Name: ", p1.first)
	fmt.Println("Last: ", p1.last)
	fmt.Println("Age :", p1.age)
}
