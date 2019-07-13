package slice

import "fmt"

// PrintInt print int slice
func PrintInt(s []int) {
	fmt.Printf("slice: len[%d], cap[%d], value %v\n", len(s), cap(s), s)
}

// MakeInt make int slice with len and cap
func MakeInt(len, cap int) []int {
	return make([]int, len, cap)
}

// IterateIntIndex iterate slice's index
func IterateIntIndex(si []int) {
	fmt.Println("iterate slice's index")
	for i := range si {
		// i is the index
		fmt.Println(i)
	}
}

// IterateInt iterate slice
func IterateInt(si []int) {
	fmt.Println("iterate slice")
	for i, elem := range si {
		fmt.Printf("[%d-th]: %v\n", i, elem)
	}
}

// Example1 shows when using slice[start:end], the underlying array is still bind together
func Example1() {
	si := []int{1, 2, 3, 4, 5}
	sub := si[1:4]
	sub[0] = 100
	PrintInt(si)
}
