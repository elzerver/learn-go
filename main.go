package main

import "fmt"

var z = 43
var f int // The zero value, assigns a 0 as a default value

func main() {
	// short declaration operator :=
	x := 42
	fmt.Println(x)
	person := "James Bond"
	fmt.Println(person)
	
	fmt.Println(z)

	foo()

	fmt.Println(f)
    // This is a static programming language
	// This statement is using the fmt.Printf not fmt.Println
	fmt.Println("This is to know the type of the variable z in this case: ")
	fmt.Printf("%T\n", z)


	// Format the text output with double quotes within a text
	// Define a variable with a string and print out with the following
	var message string = `James said, "Shaken, not stirred"`
	fmt.Println(message)
}

func foo() {
	fmt.Println("again:", z)
}