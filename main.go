package main

import "fmt"
// Importing the runtime package
import "runtime"

var z = 43
var f int // The zero value, assigns a 0 as a default value

// Define a type of data
type hotdog int 
// Asign a value using the type of hotdog
var perrito hotdog

// Define int8 value
var value1 int8
var value2 int

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


	// Learnging of the zero value
	// if you define a var without a value that returns 0 (zero)
	fmt.Println(f)
	fmt.Printf("%T\n", f)


	// Asign a value to the type of hotdog and print the output
	perrito = 23
	// Print the new type hotdog to see the value
	fmt.Println(perrito)
	// Print to see the type of the variable perrito
	fmt.Printf("%T\n", perrito)


	// Working with casting, in go is named as Conversion
	// If i try to Convert the value of perrito into z it fails
	// I need to Convert the value of perrito into int
	z = int(perrito)
	// The value of perrito is 23 ;)
	fmt.Println(perrito)
	fmt.Printf("%T\n", z)

    
    // Define the value types
	// When you define a value as the following it throws an overflow
	// on int8
    // value1 = -129
	value1 = -128
	value2 = 10

	// Printing unsigned and signed type of values
	fmt.Println(value1)
	fmt.Printf("%T\n", value1)
	fmt.Println(value2)
	fmt.Printf("%T\n", value2)

	// Using the runtime package to see the architecture
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

    // Printing the ASCII and the hexadecimal of a string 
	s := "Hellow playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Print the value of each letter
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// Iterate and print the %U# of each element
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}
}

func foo() {
	fmt.Println("again:", z)
}