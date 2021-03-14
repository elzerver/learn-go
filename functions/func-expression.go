package main

import (
	"fmt"
)

func main ()  {
	
	f := func(){
		fmt.Println("Inside a func expression")
	}
	f()

	g := func(x int){
		fmt.Println("Inside a func expression passing arguments", x)
	}
	g(1800)
}