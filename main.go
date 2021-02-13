package main

import "fmt"

var z = 43

func main() {
	// short declaration operator :=
	x := 42
	fmt.Println(x)
	person := "James Bond"
	fmt.Println(person)
	
	fmt.Println(z)

	foo()
}

func foo() {
	fmt.Println("again:", z)
}