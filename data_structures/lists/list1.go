package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Hello playground!")

	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	fmt.Println(intList.Len())

	fmt.Printf("%T\n", intList)
	// fmt.Println(*intList)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
