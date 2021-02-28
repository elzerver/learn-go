package main

import "fmt"

// Create a user defined struct with
// 	the identifier "person"
//  the fields:
//    first
//	  last
// 	  age
// attach a method to type person with
//	the identifier "speak"
//  the method should have the person say their name and age
// Create a value of type person
// Call the method from the value of type person

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hey %s, Do you really you have %d years old?\n", p.first, p.age)
}

func main() {
	fmt.Println("Hello go!")

	p1 := person{
		first: "James",
		last:  "Cameron",
		age:   27,
	}

	// fmt.Println(p1)
	p1.speak()
}
