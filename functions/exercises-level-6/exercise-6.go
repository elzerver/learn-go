package main

import "fmt"

func main() {

	defer fmt.Println("Hello wuold!")
	func() {
		fmt.Println("This is the anonymous function")
	}()

	fmt.Println("This is the second message at the end of the main")

}
