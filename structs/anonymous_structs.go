package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Peter",
		last:  "Parker",
		age:   27,
	}

	fmt.Println(p1.first)
	fmt.Println(p1)
}
