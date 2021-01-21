package maptest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	var m map[int]string
	assert.Nil(t, m)

	// nil map can't insert element
	// m[0] = "zero"
	// assert.Equal(t, "zero", m[0])

	m2 := map[int]string{
		1: "one",
		2: "two",
	}
	m2[0] = "zero"
	assert.Equal(t, "zero", m2[0])

	m3 := make(map[int]string)
	m3[0] = "zero"
	assert.Equal(t, "zero", m3[0])
}

func TestAccess(t *testing.T) {
	m := map[int]int{1: 100, 2: 200, 3: 300}

	// add element
	m[4] = 400

	// get value
	v, ok := m[4]
	assert.True(t, ok)
	assert.Equal(t, v, m[4])
	// m[4] is not modify, because v is only a copy
	v = 500
	assert.Equal(t, 400, m[4])

	v, ok = m[5] // won't create element
	assert.False(t, ok)
	assert.Len(t, m, 4)

	// delete element
	delete(m, 4)
	assert.Len(t, m, 3)

	delete(m, 5) // no-op
	m = nil
	delete(m, 5) // no-op
}

func TestStructValue(t *testing.T) {
	type People struct {
		Name string
		Age  int
	}

	m := map[string]People{"Lily": {"Lily", 29}, "SP": {"SP", 31}}
	// compiler error : cannot assign to struct field m["Lily"].Age in map
	// m["Lily"].Age = 30
	_ = m
}
