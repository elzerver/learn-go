package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {

	p.name = "Morgan"

}

func main() {

	p1 := person{
		name: "Peter",
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)

}
