package main

import "fmt"

type person struct {
	first      string
	last       string
	ice_flavor []string
}

func main() {

	p1 := person{
		first:      "Peter",
		last:       "Parker",
		ice_flavor: []string{"vanilla", "strawberry"},
	}

	p2 := person{
		first:      "Cat",
		last:       "Woman",
		ice_flavor: []string{"mice", "chocolate"},
	}

	fmt.Println(p1.first)
	for i, v := range p1.ice_flavor {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.ice_flavor)
	for i, v := range p2.ice_flavor {
		fmt.Println(i, v)
	}

	fmt.Println("Hello playground!")
}
