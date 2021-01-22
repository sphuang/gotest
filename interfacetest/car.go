package interfacetest

import "fmt"

type Car struct {}

func (c *Car) Start() {
	if c == nil {
		fmt.Println("Are you kidding?")
	} else {
		fmt.Println("Car Start")
	}
}

func (c *Car) Move() {
	fmt.Println("Car Move")
}

func (c *Car) Park() {
	fmt.Println("Car Park")
}