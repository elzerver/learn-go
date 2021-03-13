package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() *person {
	m := person{name: "noname", age: 50}
	fmt.Printf("initPerson --> %p\n", &m)
	return &m
}

func main() {
	// fmt.Println("Hello playground!")
	fmt.Printf("main --> %p\n", initPerson())
}
