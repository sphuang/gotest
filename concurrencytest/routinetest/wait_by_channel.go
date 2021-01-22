package routinetest

import (
	"fmt"
)

func RaceStartWaitByChannel(cars []*Car) {
	fmt.Println("Race Start!")

	cs := make([]chan bool, len(cars))
	for i := range cs {
		cs[i] = make(chan bool)
	}

	for i, car := range cars {
		go car.RunWithChannel(cs[i])
	}

	for _, c := range cs {
		<- c
	}

	fmt.Println("Race End!")
}