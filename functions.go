package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
	wooResponse := woo("Sting")
	fmt.Println(wooResponse)

	a, b := mouse("Richard", "Aquarius")
	fmt.Println(a)
	fmt.Println(b)

}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println(s)

}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}