package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar() {
	fmt.Println("Hello from bar")
}