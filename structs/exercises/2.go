package main

import "fmt"

type person struct {
	first      string
	last       string
	ice_flavor []string
}

func main() {

	p1 := person{
		first:      "Peter",
		last:       "Parker",
		ice_flavor: []string{"vanilla", "strawberry"},
	}

	p2 := person{
		first:      "Cat",
		last:       "Woman",
		ice_flavor: []string{"mice", "chocolate"},
	}

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m["James"])
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss moneypenny"]; ok {
		fmt.Println("This is the if print", v)
	}

	p := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(p)

	for _, v := range p {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.ice_flavor {
			fmt.Println(i, val)
		}

	}

	fmt.Println(p1.first)
	// for i, v := range p1.ice_flavor {
	// 	fmt.Println(i, v)
	// }

	fmt.Println(p2.first)
	// fmt.Println(p2.ice_flavor)
	// for i, v := range p2.ice_flavor {
	// 	fmt.Println(i, v)
	// }

	fmt.Println("Hello playground!")
}
