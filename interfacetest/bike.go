package interfacetest

import "fmt"

type Bike struct {}

func (b *Bike) Start() {
	fmt.Println("Bike Start")
}

func (b *Bike) Move() {
	fmt.Println("Bike Move")
}

func (b *Bike) Park() {
	fmt.Println("Bike Park")
}