package main

import (
	"fmt"
)

func punteros() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // Dereference the address and gives you the value stored there
}

func main() {
	punteros()
}
