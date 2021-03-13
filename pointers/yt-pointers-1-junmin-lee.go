package main

// From YT Jumin Lee tutorial
import (
	"fmt"
)

var foo int = 23

func main() {
	fmt.Println(&foo) // Address of the pointer
	fmt.Println(foo)  // Value of the variable
	fmt.Println("Hello Playground!")

	i, j := 42, 2701

	fmt.Println(i, j)
	fmt.Println(&i, &j) // Printing the address of i and j

	p := &i         // This is a pointer, pointing to the address and value of i
	fmt.Println(*p) // This is called dereferencing
	fmt.Printf("%T\n", p)

	*p = 21 // Changing the value of i
	fmt.Println(i)

	p = &j // Assign the address to p
	// It's possible because the type of p is int and the variable
	// is a type of int

	fmt.Println("The value of p :", *p)
	fmt.Printf("P: %d divided by 37 is..\n", *p)
	*p = *p / 37
	fmt.Println("Printing the new pointer")
	fmt.Println(*p)
	fmt.Println("Printing the value of j", j)

	p1 := *p
	p2 := *p
	p3 := *p

	fmt.Println(p1, p2, p3)
}
