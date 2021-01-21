package scopetest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackageVariable(t *testing.T) {
	assert.Equal(t, 1, privateA) // can access private package scope variable
	assert.Equal(t, 2, PublicA)
}

func TestPackagePrivateMethod(t *testing.T) {
	assert.Equal(t, privateA, getPrivateA()) // can access private package scope method
	assert.Equal(t, privateA, GetPrivateA())
	assert.Equal(t, PublicA, getPublicA()) // can access private package scope method
	assert.Equal(t, PublicA, GetPublicA())
}

func TestLocalVarPrioritizeGlobalVar(t *testing.T) {
	assert.Equal(t, 2, PublicA)

	if true {
		PublicA := 3
		assert.Equal(t, 3, PublicA)
	}

	assert.Equal(t, 2, PublicA)
}
