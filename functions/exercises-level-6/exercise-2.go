package main

import "fmt"

// Create a func with the idenfier foo
// takes in a variadic parameter of int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in

// create a function with the identifier bar
// that takes the parameter of type []int
// returns the sum of all values of type int passed in

func foo(values ...int) int {
	var total int
	total = 0
	for _, v := range values {
		// fmt.Println(v)
		total += v
	}
	// fmt.Println(total)
	// fmt.Printf("%T\n", total)
	return total
}

func bar(v []int) int {
	// fmt.Printf("%T\n", v)
	var total int
	total = foo(v...)
	// fmt.Printf("%T\n", total)
	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(numbers...)
	fmt.Println(n)
	var value int
	value = bar(numbers)
	fmt.Println(value)
}
