package main

import (
	"fmt"
)

// Create a type square
// Create a type circle
// attach a method to each that calculates are aand returns it
//	cicrle area = pi r 2
//  square area = L * W

// create a type shape that defines an interface as antything that has the area method
// create a func info which takes type shape and then prints the area
// create a value of type circle
// use func info to print the area of square
// use func info to print the area fo circle
const pi float64 = 3.14164

type square struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.height * s.width
}

func (c circle) area() float64 {
	return pi * (c.radius * 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	defer fmt.Println("Hello go!")

	s := square{width: 10, height: 20}
	info(s)

	c := circle{
		radius: 10,
	}
	info(c)

	fmt.Println("END")
}
