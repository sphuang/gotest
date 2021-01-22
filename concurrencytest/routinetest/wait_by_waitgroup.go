package routinetest

import (
	"fmt"
	"sync"
)

func RaceStartWaitByWaitGroup(cars []*Car) {
	fmt.Println("Race Start!")

	var wg sync.WaitGroup
	wg.Add(len(cars))

	for _, car := range cars {
		go car.RunWithWaitGroup(&wg)
	}

	wg.Wait()

	fmt.Println("Race End!")
}