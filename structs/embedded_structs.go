package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "Peter",
			last:  "Parker",
			age:   27,
		}, ltk: true,
	}

	fmt.Println(sa1)

	fmt.Println("Name: ", sa1.first)
	fmt.Println("Last: ", sa1.last)
	fmt.Println("Age :", sa1.age)
	fmt.Println("LTK :", sa1.ltk)
}
