package main

import "fmt"

func main() {

	defer fmt.Println("HELO")

	f := func() {
		fmt.Println("Hello fromt this function variable...")
	}
	f()

}
