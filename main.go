package main

import "fmt"

var z = 43
var f int // The zero value, assigns a 0 as a default value

func main() {
	// short declaration operator :=
	x := 42
	fmt.Println(x)
	person := "James Bond"
	fmt.Println(person)
	
	fmt.Println(z)

	foo()

	fmt.Println(f)
}

func foo() {
	fmt.Println("again:", z)
}