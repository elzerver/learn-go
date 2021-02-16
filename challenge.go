package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello World")
	x := 33
	for {
		if x >= 123 {
			break
		}
		// fmt.Println(x)
		fmt.Printf("%v\n\t\t%#U", x, x)
		x++
	}
}
