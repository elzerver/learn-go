package main

import (
	"fmt"
)

func main() {
	a := 4
	squareVal(a)
	squareAdd(&a)
}

func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

func squareAdd(p *int) { // This type is inputting an address instead an integer
	*p *= *p
	fmt.Println(p, *p)
}
