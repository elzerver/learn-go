package main

import (
	"fmt"
)

type Course struct {
	Name    string
	Price   float64
	isFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c Course) PrintClasses() {
	text := "Las claseses son: "
	for _, v := range c.Classes {
		fmt.Println(text, v)
	}
}

func main() {
	c1 := Course{
		Name:    "Go desde cero",
		Price:   22.13,
		isFree:  true,
		UserIDs: []uint{1, 2, 3, 4},
		Classes: map[uint]string{
			1: "Introduction",
			2: "Estructuras",
			3: "Maps",
		},
	}

	fmt.Println("Go!")
	fmt.Println(c1.Name)
	fmt.Println(c1.Price)
	c1.ChangePrice(19.99)
	fmt.Println(c1.Price)

	c1.PrintClasses()
}
