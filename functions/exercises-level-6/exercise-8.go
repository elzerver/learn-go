package main

import "fmt"

func message() func() int {
	return func() int {
		return 43
	}
}

func main() {
	defer fmt.Println("Hello playground!")

	f := message()
	fmt.Println(f())
}
