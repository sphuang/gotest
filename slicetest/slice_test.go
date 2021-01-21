package slicetest

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Car struct {
	Price int
	Brand string
}

func TestSliceRange(t *testing.T) {
	cars := []Car{ Car{50, "toyota"}, Car{300, "bmw"}, Car{500, "porsche"} }
	for _, car := range cars {
		car.Price += 10
	}

	// below tests will be failed, because car is only the copy of cars[i]
	// assert.Equal(t, 60, cars[0].Price)
	// assert.Equal(t, 310, cars[1].Price)
	// assert.Equal(t, 510, cars[2].Price)
}

func TestSliceRange2(t *testing.T) {
	cars := []Car{ Car{50, "toyota"}, Car{300, "bmw"}, Car{500, "porsche"} }
	for i, _ := range cars {
		cars[i].Price += 10
	}
	assert.Equal(t, 60, cars[0].Price)
	assert.Equal(t, 310, cars[1].Price)
	assert.Equal(t, 510, cars[2].Price)
}

func TestSliceAssignment(t *testing.T) {
	names := []string{"Lily", "SP", "Apple"}
	names2 := names

	// slice assignment will share the same underlying array
	assert.Equal(t, &names[0], &names2[0])

	// slice assignment will share the same underlying array
	var names3 []string
	names3 = names
	assert.Equal(t, &names[0], &names3[0])
}

func TestAppend(t *testing.T) {
	scores := []int{1, 2, 3, 4, 5}
	log.Printf("cap = %d", cap(scores))
	count := 0
	for {
		addr := &scores
		maxLen := cap(scores)
		scores = append(scores, 6)
		newAddr := &scores
		newMaxLen := cap(scores)
		if addr != newAddr {
			log.Printf("different at %d", len(scores))
		}
		if maxLen != newMaxLen {
			log.Printf("different cap [ori = %d, new = %d]", maxLen, newMaxLen)
		}
		count++
		if count > 100000 {
			log.Printf("no meet at length %d", len(scores))
			break
		}
	}
	assert.False(t, true)
}

func TestInit(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 5, len(a))
	assert.Equal(t, 5, cap(a))

	b := make([]int, 3)
	assert.Equal(t, 3, len(b))
	assert.Equal(t, 3, cap(b))

	c := make([]int, 3, 5)
	assert.Equal(t, 3, len(c))
	assert.Equal(t, 5, cap(c))

	var d []int
	assert.Nil(t, d)
	assert.Zero(t, len(d))
	assert.Zero(t, cap(d))

	e := []int{}
	assert.False(t, e == nil)
	assert.Zero(t, len(e))
	assert.Zero(t, cap(e))
}

func TestAccess(t *testing.T) {
	a := make([]int, 3, 5)
	a[2] = 3
	// a[3] = 4 // panic, out-of-bound
}

func TestSubslice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// share the same underlying array
	b := a[:5]
	b[0] = 6
	assert.Equal(t, b[0], a[0])
	assert.Equal(t, 5, len(b))
	assert.Equal(t, 10, cap(b))

	// share the same underlying array
	c := a[5:]
	c[0] = 10
	assert.Equal(t, c[0], a[5])
	assert.Equal(t, 5, len(c))
	assert.Equal(t, 5, cap(c))
}

func Test2DSliceAppend(t *testing.T) {
	var s2D [][]int
	s := []int{1, 2, 3}
	s2D = append(s2D, s)

	s = s[:2]
	s = append(s, 4)
	s2D = append(s2D, s)

	// will print [[1 2 4] [1 2 4]]  ---> be careful !!, the underlying array still connect together
	log.Println(s2D)
}
