package main

import (
	"fmt"
)

func main() {
	// This code does not run
	c := make(chan int, 1)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
