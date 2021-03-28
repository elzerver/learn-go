package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	car       vehicle
	fourWheel bool
}

type sedan struct {
	car    vehicle
	luxury bool
}

func main() {
	fmt.Println("Creating two vehicles")

	monsterTruck := truck{
		car: vehicle{
			doors: "4",
			color: "blue",
		},
		fourWheel: false,
	}

	superCar := sedan{
		car: vehicle{
			doors: "4",
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(monsterTruck.car)
	fmt.Println(superCar.car.doors)
}
