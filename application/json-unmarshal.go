package main

import (
	"encoding/json"
	"fmt"
)

// This needs to create the structure data type in order to parse the json data
// Remember that this is a slice of struct that returns the data.
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	defer fmt.Println("Hello playground")

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"MoneyPenny","Age":27}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data", people)

	for i, v := range people {
		fmt.Println("PERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
