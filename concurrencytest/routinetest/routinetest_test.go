package routinetest

import (
	// "log"
	// "reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoWait(t *testing.T) {
	RaceStart([]*Car{{"Jonny"}, {"Tony"}, {"Jimmy"}})
}

func TestWaitByChannel(t *testing.T) {
	RaceStartWaitByChannel([]*Car{{"Jonny"}, {"Tony"}, {"Jimmy"}})
}

func TestWaitByWaitGroup(t *testing.T) {
	RaceStartWaitByWaitGroup([]*Car{{"Jonny"}, {"Tony"}, {"Jimmy"}})
}

func TestShareResource(t *testing.T) {
	assert.NotEqual(t, 10000, ConcurrentPlusFund(false))
	assert.Equal(t, 10000, ConcurrentPlusFund(true))
}
