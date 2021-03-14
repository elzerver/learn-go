package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Last string
	Age  int
}

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (p person) Speak() {
	fmt.Printf("Hi %s %s , are you %d years old?\n", p.Name, p.Last, p.Age)
}

func main() {

	p1 := person{
		Name: "Peter",
		Last: "Parker",
		Age:  27,
	}

	p2 := person{"Jorge", "Morales", 30}
	p3 := person{"Jaime", "Duende", 27}

	people := []person{p1, p2, p3}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

	fmt.Println(p1)
	p1.Speak()
	fmt.Println("Hello playground")
}
