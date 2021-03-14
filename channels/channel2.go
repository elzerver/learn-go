package main

import (
	"fmt"
)

func main() {
	// This code does not run
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
