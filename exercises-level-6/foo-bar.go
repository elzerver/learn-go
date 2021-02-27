package main

import (
	"fmt"
)

func foo() int {
	x := 10
	return x
}

func bar() (int, string) {
	x := 5
	y := "I'm a string!"

	return x, y
}

func main() {
	msg := foo()
	fmt.Println("Give me a number: ", msg)

	i, j := bar()

	fmt.Println("This are the values", i, j)
}
