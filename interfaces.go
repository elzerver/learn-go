package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "- The secret Agent speak")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "- The person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}
func main() {
	sa1 := secretAgent{
		person: person{
			first: "Rambo",
			last:  "Perez",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Peter",
		last:  "Parker",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
