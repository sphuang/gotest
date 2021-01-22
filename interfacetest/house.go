package interfacetest

import "fmt"

type House struct {}

func (h *House) Enter() {
	fmt.Println("Enter House")
}

func (h *House) Leave() {
	fmt.Println("Leave House")
}
