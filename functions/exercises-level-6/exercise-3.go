package main

import (
	"fmt"
)

func foo() {
	defer func() {
		fmt.Println("foo defer ran")
	}()

}

func bar() {
	fmt.Println("In bar")
}

func main() {

	defer foo()
	bar()

	foo()
	bar()

}
