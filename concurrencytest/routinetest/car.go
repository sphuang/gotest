package routinetest

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

type Car struct {
	Name string
}

func (c *Car) Run() {
	time.Sleep(time.Second * time.Duration(rand.Float64() * 5))
	fmt.Printf("Car [%v] has finished!\n", c.Name)
}

func (c *Car) RunWithChannel(finish chan<- bool) {
	c.Run()
	finish <- true
}

func (c *Car) RunWithWaitGroup(wg *sync.WaitGroup) {
	c.Run()
	wg.Done()
}