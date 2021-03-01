package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "MoneyPenny",
		Age:   27,
	}

	people := []person{p1, p2}

	// fmt.Println(people)

	fmt.Println("Hello world!")
	bs, err := json.Marshal(people)

	// fmt.Println(bs, err)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
