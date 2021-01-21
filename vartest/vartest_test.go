package vartest

import (
	"github.com/stretchr/testify/assert"
	"testing"
	// "fmt"
)

var globalI int
var globalF float64

func returnLocalVarPointer() *int {
	a := 3
	return &a
}

// It's OK to return pointer of local variable, because of escape analysis
func TestEscapeAnalysis(t *testing.T) {
	a := returnLocalVarPointer()
	assert.Equal(t, 3, *a)
}

func TestDefaultInitToZero(t *testing.T) {
	// local variable: default to zero value
	var b bool
	var i int
	var f float64
	var s string
	var m map[int]string
	var iface interface{}
	var sl []int

	assert.Equal(t, false, b)
	assert.Equal(t, 0, i)
	assert.Equal(t, 0.0, f)
	assert.Equal(t, "", s)
	assert.Nil(t, m)
	assert.Nil(t, iface)
	assert.Nil(t, sl)

	// struct member: default to zero value
	type Car struct {
		Price int
		Brand string
	}
	var car Car

	assert.Equal(t, 0, car.Price)
	assert.Equal(t, "", car.Brand)

	// global variable: default to zero value
	assert.Equal(t, 0, globalI)
	assert.Equal(t, 0.0, globalF)
}
