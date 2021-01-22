package routinetest

import (
	"fmt"
)

func RaceStart(cars []*Car) {
	fmt.Println("Race Start!")
	for _, car := range cars {
		go car.Run()
	}
	fmt.Println("Race End!")
}