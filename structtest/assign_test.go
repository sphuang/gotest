package structtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignment(t *testing.T) {
	house1 := MyHouse{"An-Le street", 3, []string{"Lily", "SP", "Apple"}}

	// struct assignment is shallow copy
	house2 := house1

	assert.Equal(t, house1.Address, house2.Address)
	assert.Equal(t, house1.PeopleNum, house2.PeopleNum)
	assert.Equal(t, house1.Peoples, house2.Peoples)
	assert.Equal(t, &house1.Peoples[0], &house2.Peoples[0]) // shallow copy: underlying array are the same
}
